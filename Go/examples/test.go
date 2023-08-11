package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Apps map[string]string `yaml:""`
}

func main() {
	// Read the file
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a struct to hold the YAML data
	var config Config

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the data
	fmt.Println(config.Name)
}
