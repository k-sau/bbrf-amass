package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/k-sau/bbrf-amass/pkg/bbrf"
	"github.com/k-sau/bbrf-amass/pkg/core"
)

func main() {
	filePath := flag.String("path", "", "Full path to amass json output.")
	program := flag.String("p", "", "Program id. Required.")
	help := flag.Bool("h", false, "Prints available flags")
	service := flag.Bool("service", false, "Takes input from stdin in format of ip;port;service-name. Supports ipv6")
	bbrfConfigFile := flag.String("bc", "~/.bbrf/config.json", "File path for bbrf config file. Default: ~/.bbrf/config.json")
	wildcard := flag.Bool("wildcard", false, "Adds everything excepts domains which explicitly mentioned in out of scope.")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(1)
	}
	if *program == "" {
		fmt.Println("p is required")
		flag.Usage()
		os.Exit(1)
	}
	if !*service && *filePath == "" {
		fmt.Println("service or path is required")
		flag.Usage()
		os.Exit(1)
	}

	if *service && *filePath != "" {
		fmt.Println(" Any one flag is required: service or path")
		flag.Usage()
		os.Exit(1)
	}
	if *service && *filePath != "" {
		fmt.Println("any one flag is required: service or path")
		flag.Usage()
		os.Exit(1)
	}
	// Parse config file
	core.ParseConfigFile(*bbrfConfigFile)

	if *service {
		bbrf.Initialize("service", *program)
	} else if *filePath != "" {
		core.Initialize(*filePath, *program, *wildcard)
	}
}
