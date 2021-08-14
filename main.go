package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/k-sau/bbrf-amass/pkg/core"
)

func main() {
	filePath := flag.String("path", "", "Full path to amass json output. Required.")
	program := flag.String("p", "", "Program id. Required.")
	help := flag.Bool("h", false, "Prints available flags")
	bbrfConfigFile := flag.String("bc", "~/.bbrf/config.json", "File path for bbrf config file. Default: ~/.bbrf/config.json")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(1)
	}

	if *program == "" || *filePath == "" {
		fmt.Println("path and p is required")
		flag.Usage()
		os.Exit(1)
	}
	core.Initialize(*filePath, *program, *bbrfConfigFile)
}
