package bbrf

import (
	"bufio"
	"os"
	"strings"

	"github.com/k-sau/bbrf-amass/pkg/constants"
	"github.com/k-sau/bbrf-amass/pkg/core"
)

// Handles domain names from stdin

func unresolvedInitialize(program string) {
	var docs constants.UnResolvedDocuments
	// read the lines, append and output them if they're new
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := sc.Text()

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
