package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
)

func main() {

	configPath := flag.String("config", "", "Configuration file path")
	flag.Parse()
	if *configPath == "" {
		panic("ERROR: config Path is required")
	}

	rawData, err := ioutil.ReadFile(*configPath)
	if err != nil {
		panic("ERROR: error while reading the file")
	}
	// convert yaml to json
	yamlToJson(rawData)
}

func yamlToJson(rawData []byte) {
	var body interface{}
	if err := yaml.Unmarshal(rawData, &body); err != nil {
		panic(err)
	}

	body = convert(body)

	if b, err := json.Marshal(body); err != nil {
		panic(err)
	} else {
		fmt.Printf("Output: %s\n", b)
	}
}

// unmarshal into a value of type interface{}, and go through the result recursively,
// and convert each encountered map[interface{}]interface{} to a map[string]interface{} value
func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}
