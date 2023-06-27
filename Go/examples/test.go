package main

import (
	"encoding/json"
	"fmt"
)

type Target struct {
	Labels  map[string]string `json:"labels"`
	Targets []string          `json:"targets"`
}

type Result struct {
	Error  string   `json:"error"`
	Result []string `json:"result"`
}

func main() {
	jsonData := []byte(`{
		"error": "",
		"result": [
			"gatwickexpress.on.icomera.com",
			"https://greateranglia-stg.on.icomera.com",
			"aaa.on.icomera.com",
			"https://google.co.uk"
		]
	}`)

	var data Result
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	targets := Target{
		Labels: map[string]string{
			"app": "cna",
		},
		Targets: data.Result,
	}

	output, err := json.Marshal([]Target{targets})
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(string(output))
}
