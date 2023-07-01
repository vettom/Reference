package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type Target struct {
	Labels  map[string]string `json:"labels"`
	Targets []string          `json:"targets"`
}

type Result struct {
	Result []string `json:"result"`
}

func fetchData() {
	url := "https://vettom.github.io/api/cna.json"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error: failed to get URL \n", err)
		return
	}
	defer response.Body.Close()
	// Read body of URL
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Failed to read response data! \n", err)
		return
	}

	// Unmarshal data
	var jsonData Result
	// err := json.Unmarshal(data, &jsonData)
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Fatal("Error decoding JSON: \n", err)
		return
	}
	fmt.Println(jsonData)
}

func fetchEndpoint(appname, conf string) string {
	vi := viper.New()
	vi.SetConfigFile(conf)
	vi.ReadInConfig
	if err != nil {
		log.Fatal("Error: failed to open endpoints.yaml")
	}

}

func main() {
	appname := "cna"
	conf := "endpoints.yaml"
	link := fetchEndpoint(appname, conf)
	fmt.Println(link)
	// fetchData()
}
