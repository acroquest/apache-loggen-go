package main

import (
	"github.com/acroquest/apache-loggen-go"
)

var config loggen.Config

func init() {
	config = loggen.LoadConfig("./config.json")
}

func main() {
	loggen.GenerateNewRecord(config)
}
