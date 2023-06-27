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
	Result []string `json:"result"`
}

func main() {
	// appName := "cna"
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

	output := Target{
		Labels: map[string]string{
			"application": "cna",
		},
		Targets: data.Result,
	}
	jsonOut, err := json.Marshal([]Target{output})
	if err != nil {
		fmt.Println("Fail")
	}
	fmt.Println(string(jsonOut))

}
