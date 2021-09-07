package bbrf

import (
	"bufio"
	"os"
	"strings"

	"github.com/k-sau/bbrf-amass/pkg/constants"
	"github.com/k-sau/bbrf-amass/pkg/core"
)

func serviceInitialize(program string) {
	var obj constants.ServiceDocuments
	var doc constants.ServiceDocument
	var keys core.Keys
	var keyArr []string

	// read the lines, append and output them if they're new
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		doc.Tags = make(map[string]string)
		service := sc.Text()
		// Uses: https://github.com/k-sau/nmap-parse-output/blob/master/nmap-parse-output-xslt/bbrf-service.xslt
		// for parsing nmap xml output
		temp := strings.Split(service, ";")
		doc.Id = temp[0] + ":" + temp[1]
		doc.Program = program
		doc.Type = "service"
		doc.Ip = temp[0]
		doc.Port = temp[1]
		doc.Service = ""
		if len(temp) == 3 {
			doc.Service = temp[2]
			doc.Tags[temp[2]] = "true"
		}
		obj.Docs = append(obj.Docs, doc)
		keyArr = append(keyArr, doc.Id)
	}

	//b, err := json.Marshal(obj)
	//core.HandleError(err)

	//fmt.Println(string(b))
	core.AddDataToBBRF(obj)
	keys.Key = keyArr
	currentDocs := core.GetCurrentBBRFData(keys, "service")
	core.UpdateBBRFData(currentDocs, obj)
}
