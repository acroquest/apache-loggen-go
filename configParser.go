package loggen

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Filename     string  `json:"filename"`
	Prefix       string  `json:"prefix"`
	Days         int     `json:"days"`
	NumOfFiles   int     `json:"num_of_files"`
	ErrRate      float64 `json:"error_rate"`
	Bytes        int     `json:"bytes"`
	ResponseTime int     `json:"response_time"`
}

func LoadConfig(filename string) Config {
	var config Config

	configFile, err := os.Open(filename)
	if err != nil {
		log.Fatal("Opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		log.Fatal("Parsing config file", err.Error())
	}

	return config
}
