package main

import (
	"flag"
)

// ConfSpec type for global config.
type ConfSpec struct {
	YamlForParse string
	Debug        bool
}

// Configuration args.
var globalConfig ConfSpec

func globalConfigInit() {
	// Read flags.
	// Read yaml file name option -f ( -f filename.yaml )
	flag.StringVar(&globalConfig.YamlForParse, "f", "", "Yaml file.")
	// Read debug option ( --debug )
	flag.BoolVar(&globalConfig.Debug, "debug", false, "Turn on debug logging. Default: false")
	// Parse args.
	flag.Parse()
	// Check if filename is set.
	if globalConfig.YamlForParse == "" {
		log.Fatalf("Option -f is required.")
	}
}
