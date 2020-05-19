package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Recursive iterate through parsed yaml, print variables definitions.
func iterateAndPrintStruct(value interface{}, prefix string) {
	if value == nil {
		return
	}
	// Separator for variable name.
	sp := "_"
	if prefix == "" {
		// This is a iteration with recursive depth 1
		// Use empty separator.
		sp = ""
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice:
		// Value type is slice (array) - need recursive call. Add separator and index as a prefix.
		for index, v := range value.([]interface{}) {
			iterateAndPrintStruct(v, prefix+sp+strconv.Itoa(index))
		}
	case reflect.Map:
		// Value type is map - need recursive call. Add separator and key name as a prefix.
		for name, v := range value.(map[interface{}]interface{}) {
			iterateAndPrintStruct(v, prefix+sp+name.(string))
		}
	default:
		// Value is value :) (string, int, bool, float64 ...). Print variable definition for bash.
		fmt.Printf("%s=%v\n", prefix, value)
	}
}

func main() {
	// Init config (args)
	globalConfigInit()
	loggingInit()
	// Read main yaml file.
	mainYaml, err := ioutil.ReadFile(globalConfig.YamlForParse)
	if err != nil {
		log.Fatalf("Error reading yaml file: %s", err.Error())
	}
	// Parse yaml file.
	var parsedYaml interface{}
	err = yaml.Unmarshal([]byte(mainYaml), &parsedYaml)
	if err != nil {
		log.Fatalf("Error parsing yaml file: %s", err.Error())
	}
	// Iterate interface structure and output variables definitions for bash source <(yamltoenv).
	iterateAndPrintStruct(parsedYaml, "")

}
