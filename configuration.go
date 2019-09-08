package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration is the top most level of the config
type Configuration struct {
	Config ConfigData
}

// ConfigData is the individual configuration
type ConfigData struct {
	DiffCommand       string
	FileCompareFolder string
	FilePrefix        string
}

func getConfig() Configuration {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}
