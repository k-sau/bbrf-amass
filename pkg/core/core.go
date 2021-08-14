package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func Initialize(filepath, program, bbrfClient string) {
	handleJSONOutput(filepath, program, bbrfClient)
}

func handleJSONOutput(filepath, program, bbrfClient string) {
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

		re := regexp.MustCompile(`:.*:.*`)
		commandInjection := regexp.MustCompile(`[^A-Za-z0-9\._\-:]`)
		var domain []string

		//fmt.Println(len(obj.Addresses))
		if len(obj.Addresses) == 0 {
			domain = append(domain, obj.Name)
		} else {
			if commandInjection.MatchString(obj.Name) {
				fmt.Println("Some wrong with this domain. Skipping this one since this matches with this [^A-Za-z0-9\\._\\-:]", obj.Name)
				continue
			}
			for _, v := range obj.Addresses {
				if commandInjection.MatchString(v.Ip) {
					fmt.Println("Some wrong with this IP. Skipping this one since this matches with this [^A-Za-z0-9\\._\\-:]", v.Ip)
					continue
				}
				temp := obj.Name
				temp += ":"
				temp += v.Ip
				//fmt.Println("Before", temp)
				temp = re.ReplaceAllString(temp, "")

				//fmt.Println(temp)
				domain = append(domain, temp)
			}
		}

		sources := ""
		for _, v := range obj.Sources {
			sources += v
			sources += ","
		}
		sources = sources[:len(sources)-1]
		//fmt.Println(domain, sources)
		addDataToBBRF(domain, sources, program, bbrfClient)
	}
}

func addDataToBBRF(domains []string, sources, program, bbrfClient string) {
	for _, v := range domains {
		// command := fmt.Sprintf("/home/op/.local/bin/bbrf domain add %s -p %s -s %s", v, program, sources)
		//fmt.Println(command)
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		command := fmt.Sprintf("%s domain add %s -p %s -s %s", bbrfClient, v, program, sources)
		cmd := exec.Command("bash", "-c", command)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()
		if stderr.String() != "" {
			fmt.Println(stderr.String())
		}
		handleError(err)
	}
}
