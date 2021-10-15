package bbrf

import (
	"bufio"
	"os"
	"strings"

	"github.com/k-sau/bbrf-amass/pkg/constants"
	"github.com/k-sau/bbrf-amass/pkg/core"
)

// Handles domain names from stdin

func unresolvedInitialize(program string, wildcard bool) {
	var docs constants.UnResolvedDocuments

	// Get scopes
	// 1 = inscope wildcard
	// 2 = inscope fqdn

	// read the lines, append and output them if they're new
	sc := bufio.NewScanner(os.Stdin)
	inscope, outscope := core.GetScope(program)

	for sc.Scan() {
		domain := sc.Text()
		flag := 0

		// Skip inscope check if wildcard flag is provided
		if !wildcard {
			if core.CheckScope(domain, inscope, outscope) {
				flag = 1
			}
		} else {
			flag = 1
		}
		// If not in scope, continue.
		if flag == 0 {
			continue
		}

		// underscore is reserved keyword in couchdb
		if strings.HasPrefix(domain, "_") {
			domain = "." + domain
		}

		var tmpDoc constants.UnResolvedDocument
		tmpDoc.Id = domain
		tmpDoc.Type = "domain"
		tmpDoc.Program = program
		tmpDoc.IPs = []string{}
		docs.Docs = append(docs.Docs, tmpDoc)
	}

	core.AddDataToBBRF(docs)

}
