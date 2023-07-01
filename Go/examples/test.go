package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	vi := viper.New()
	vi.SetConfigFile("endpoints.yaml")
	vi.ReadInConfig()
	for Link in v1.GetString() {
		fmt.Println(Link)
	}

}
