package main

import (
	// "flag"
	// "strconv"
	// "strings"

	"github.com/acroquest/apache-loggen-go"
)

var config loggen.Config

func init() {
	config = loggen.LoadConfig("./config.json")
}

func main() {
	// flag.Parse()

	loggen.GenerateNewRecord(config)

	/*
		if config.Filename == "" {
			loggen.GenerateLog(config)
		} else {
			// ファイル名がある場合
			splitted := strings.Split(config.Filename, ".")
			if len(splitted) == 1 {
				for i := 1; i <= config.NumOfFiles; i++ {
					// f := config.Filename + "-" + strconv.Itoa(i)
					loggen.GenerateLogToFile(config)
				}
			} else {
				// head := strings.Join(splitted[:len(splitted)-1], "")
				for i := 1; i <= config.NumOfFiles; i++ {
					// f := head + "-" + strconv.Itoa(i) + "." + splitted[len(splitted)-1]
					loggen.GenerateLogToFile(config)
				}
			}
		}
	*/
}
