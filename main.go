package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/k-sau/bbrf-amass/pkg/core"
)

func main() {
	filePath := flag.String("path", "", "Full path to amass json output")
	program := flag.String("p", "", "Program id")
	bbrfClient := flag.String("bbrf", "/home/op/.local/bin/bbrf", "Path to bbrf. Default: /home/op/.local/bin/bbrf")

	flag.Parse()

	if *program == "" || *filePath == "" {
		fmt.Println("Flag path and p is required")
		flag.Usage()
		os.Exit(1)
	}
	core.Initialize(*filePath, *program, *bbrfClient)
}
