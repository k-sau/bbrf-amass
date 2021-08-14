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
)

func Initialize(filepath, program, bbrfConfigFile string) {
	parseConfigFile(bbrfConfigFile)
	handleJSONOutput(filepath, program)
}

func parseConfigFile(bbrfConfigFile string) {
	if bbrfConfigFile == "~/.bbrf/config.json" {
		home, err := os.UserHomeDir()
		handleError(err)
		bbrfConfigFile = home + "/.bbrf/config.json"
	}
	f, err := os.Open(bbrfConfigFile)
	handleError(err)
	defer f.Close()
	byteValue, _ := ioutil.ReadAll(f)
	err = json.Unmarshal(byteValue, &conf)
	handleError(err)
}

func handleJSONOutput(filepath, program string) {

	var docs []Document
	var obj Documents
	var keys Keys
	var keyArr []string

	f, err := os.Open(filepath)
	handleError(err)
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

		sources := ""
		for _, v := range obj.Sources {
			sources += v
			sources += ","
		}
		sources = sources[:len(sources)-1]
		//fmt.Println(domain, sources)

		var ips []string
		for _, v := range obj.Addresses {
			ips = append(ips, v.Ip)
		}

		var doc Document
		doc.Id = obj.Name
		doc.Program = program
		doc.Type = "domain"
		doc.Ips = ips
		doc.Source = sources
		docs = append(docs, doc)

		keyArr = append(keyArr, obj.Name)
	}

	obj.Docs = docs
	addDataToBBRF(obj)

	keys.Key = keyArr
	getCurrentBBRFData(keys)
}

func addDataToBBRF(obj Documents) {
	b, err := json.Marshal(obj)
	handleError(err)

	client := &http.Client{}
	req, err := http.NewRequest("POST", conf.CouchDB+"/_bulk_docs", bytes.NewBuffer(b))
	handleError(err)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(conf.Username, conf.Password)
	resp, err := client.Do(req)
	handleError(err)

	if resp.StatusCode != 201 {
		fmt.Println("Something went wrong. Please check")
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		handleError(err)
		fmt.Println(string(bodyBytes))
	}
}

func updateBBRFData(obj Documents) {

}

func getCurrentBBRFData(keys Keys) {

}
