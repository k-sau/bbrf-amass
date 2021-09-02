package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/k-sau/bbrf-amass/pkg/constants"
)

func Initialize(filepath, program string) {
	inscope, outscope := getScope(program)
	handleJSONOutput(filepath, program, inscope, outscope)
}

func getScope(program string) (map[string]int, map[string]int) {
	// 1 = inscope wildcard
	// 2 = inscope fqdn
	inscope := make(map[string]int)
	outscope := make(map[string]int)

	var keys Keys
	keys.Key = append(keys.Key, program)

	b, err := json.Marshal(keys)
	HandleError(err)

	client := &http.Client{}
	req, err := http.NewRequest("POST", conf.CouchDB+"/_all_docs?include_docs=true", bytes.NewBuffer(b))
	HandleError(err)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(conf.Username, conf.Password)
	resp, err := client.Do(req)
	HandleError(err)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err)

	if resp.StatusCode != 200 {
		fmt.Println("Something went wrong. Please check")
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		HandleError(err)
		fmt.Println(string(bodyBytes))
		os.Exit(1)
	}
	var scope Scopes
	err = json.Unmarshal(bodyBytes, &scope)

	for _, v := range scope.Rows {
		for _, vv := range v.Doc.Inscope {
			if strings.HasPrefix(vv, "*.") {
				inscope[vv[1:]] = 1
			} else {
				inscope[vv] = 2
			}
		}
		for _, vv := range v.Doc.Outscope {
			outscope[vv] = 1
		}
	}

	return inscope, outscope
}
func ParseConfigFile(bbrfConfigFile string) {
	if bbrfConfigFile == "~/.bbrf/config.json" {
		home, err := os.UserHomeDir()
		HandleError(err)
		bbrfConfigFile = home + "/.bbrf/config.json"
	}
	f, err := os.Open(bbrfConfigFile)
	HandleError(err)
	defer f.Close()
	byteValue, _ := ioutil.ReadAll(f)
	err = json.Unmarshal(byteValue, &conf)
	HandleError(err)
}

func handleJSONOutput(filepath, program string, inscope, outscope map[string]int) {

	var docs []Document
	var obj Documents
	var keys Keys
	var keyArr []string

	f, err := os.Open(filepath)
	HandleError(err)
	defer f.Close()
	byteValue, _ := ioutil.ReadAll(f)

	dec := json.NewDecoder(strings.NewReader(string(byteValue)))
	for {
		var obj Output
		if err := dec.Decode(&obj); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(obj.Name, " ", obj.Addresses[0].Ip, obj.Sources)

		// Check if domain name is in scope
		flag := 0
		for i, _ := range inscope {
			if strings.HasSuffix(obj.Name, i) {
				flag = 1
				break
			}
		}

		if flag == 0 {
			continue
		}
		// Check if particular domain name is out of scope
		for i, _ := range outscope {
			if obj.Name == i {
				flag = 0
			}
		}

		if flag == 0 {
			continue
		}

		sources := sourcesToString(obj.Sources)

		//fmt.Println(domain, sources)

		var ips []string
		for _, v := range obj.Addresses {
			ips = append(ips, v.Ip)
		}

		var doc Document

		// underscore is reserved keyword in couchdb
		if strings.HasPrefix(obj.Name, "_") {
			obj.Name = "." + obj.Name
		}

		doc.Id = obj.Name
		doc.Program = program
		doc.Type = "domain"
		doc.Ips = ips
		doc.Source = sources
		//fmt.Println(doc)
		docs = append(docs, doc)

		keyArr = append(keyArr, obj.Name)
	}

	//fmt.Println(docs)
	obj.Docs = docs
	AddDataToBBRF(obj)

	keys.Key = keyArr
	currentDocs := GetCurrentBBRFData(keys, "domain")
	//fmt.Println(currentDocs)
	UpdateBBRFData(currentDocs, obj)
}

func sourcesToString(sources []string) string {
	str := ""
	for _, v := range sources {
		str += v
		str += ","
	}
	str = str[:len(str)-1]
	return str
}

func AddDataToBBRF(obj interface{}) {
	b, err := json.Marshal(obj)
	HandleError(err)

	client := &http.Client{}
	req, err := http.NewRequest("POST", conf.CouchDB+"/_bulk_docs", bytes.NewBuffer(b))
	HandleError(err)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(conf.Username, conf.Password)
	resp, err := client.Do(req)
	HandleError(err)

	if resp.StatusCode != 201 {
		fmt.Println("Something went wrong. Please check")
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		HandleError(err)
		fmt.Println(string(bodyBytes))
	}
}

func mergeDocuments(currentDocs_ interface{}, data interface{}) interface{} {

	switch data.(type) {
	case Documents:
		currentDocs := currentDocs_.(CurrentDocuments)
		// Domains merging
		// Quick mapping for better searching
		docs := data.(Documents)
		mDocs := make(map[string]Document)
		for _, v := range docs.Docs {
			mDocs[v.Id] = v
		}

		// We only need to merge IP and sources field
		var updatedCurrentDocs CurrentDocuments
		for _, v := range currentDocs.Rows {
			mmIP := make(map[string]bool)
			for _, ip := range v.UpdateDocs.Ips {
				mmIP[ip] = true
			}

			sources := strings.Split(v.UpdateDocs.Source, ",")
			mmSource := make(map[string]bool)
			for _, s := range sources {
				mmSource[s] = true
			}

			if v.UpdateDocs.Id == "" {
				continue
			}
			//fmt.Println(v.UpdateDocs.Id)
			temp := mDocs[v.UpdateDocs.Id]
			for _, vv := range temp.Ips {
				if mmIP[vv] == false {
					v.UpdateDocs.Ips = append(v.UpdateDocs.Ips, vv)
					mmIP[vv] = true
				}
			}

			tempSources := strings.Split(temp.Source, ",")

			for _, vv := range tempSources {
				if mmSource[vv] == false {
					sources = append(sources, vv)
					mmSource[vv] = true
				}
			}
			v.UpdateDocs.Source = sourcesToString(sources)
			updatedCurrentDocs.Rows = append(updatedCurrentDocs.Rows, v)
		}
		return updatedCurrentDocs

	case constants.ServiceDocuments:
		// Services merging
		docs := data.(constants.ServiceDocuments)
		// Quick mapping for better searching
		mDocs := make(map[string]constants.ServiceDocument)
		for _, v := range docs.Docs {
			mDocs[v.Id] = v
		}
		fmt.Println(mDocs)
		currentDocs := currentDocs_.(constants.ServiceCurrentDocuments)

		// We only need to update the service name
		var updatedCurrentDocs constants.ServiceCurrentDocuments
		for _, v := range currentDocs.Rows {
			v.UpdateDocs.Service = mDocs[v.UpdateDocs.Id].Service
			updatedCurrentDocs.Rows = append(updatedCurrentDocs.Rows, v)
		}

		return updatedCurrentDocs
	}
	return 1
}

func UpdateBBRFData(currentDocs interface{}, docs interface{}) {
	// Merge both documents object
	tmp := mergeDocuments(currentDocs, docs)

	var b []byte
	var err error
	switch tmp.(type) {
	case CurrentDocuments:
		//fmt.Println(updatedDocs)
		updatedDocs := tmp.(CurrentDocuments)
		var obj BulkUpdate
		for _, v := range updatedDocs.Rows {
			obj.Docs = append(obj.Docs, v.UpdateDocs)
		}

		b, err = json.Marshal(obj)
		HandleError(err)
		//fmt.Println(string(b))

	case constants.ServiceCurrentDocuments:
		updatedDocs := tmp.(constants.ServiceCurrentDocuments)
		var obj constants.ServiceBulkUpdate
		for _, v := range updatedDocs.Rows {
			obj.Docs = append(obj.Docs, v.UpdateDocs)
		}
		b, err = json.Marshal(obj)
		HandleError(err)

	case int:
		fmt.Printf("CurrentDocuments returned with status 1 insted of object. Quiting")
		os.Exit(1)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", conf.CouchDB+"/_bulk_docs", bytes.NewBuffer(b))
	HandleError(err)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(conf.Username, conf.Password)
	resp, err := client.Do(req)
	HandleError(err)

	if resp.StatusCode != 201 {
		fmt.Println("Something went wrong. Please check")
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		HandleError(err)
		fmt.Println(string(bodyBytes))
	}
}

func GetCurrentBBRFData(keys Keys, type_ string) interface{} {
	b, err := json.Marshal(keys)
	HandleError(err)

	client := &http.Client{}
	req, err := http.NewRequest("POST", conf.CouchDB+"/_all_docs?include_docs=true", bytes.NewBuffer(b))
	HandleError(err)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(conf.Username, conf.Password)
	resp, err := client.Do(req)
	HandleError(err)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err)

	if resp.StatusCode != 200 {
		fmt.Println("Something went wrong. Please check")
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		HandleError(err)
		fmt.Println(string(bodyBytes))
		os.Exit(1)
	}

	if type_ == "domain" {
		var currentDocs CurrentDocuments
		err = json.Unmarshal(bodyBytes, &currentDocs)
		HandleError(err)
		return currentDocs
	} else if type_ == "service" {
		var currentDocs constants.ServiceCurrentDocuments
		err = json.Unmarshal(bodyBytes, &currentDocs)
		HandleError(err)
		return currentDocs
	}

	return 1
}
