package main

import (
	"flag"
	"github.com/acroquest/apache-loggen-go"
)

func main() {
	var (
		days     int
		errRate  float64
		filename string
	)
	flag.IntVar(&days, "d", 1, "same as -day")
	flag.Float64Var(&errRate, "e", 0.1, "same as -err")
	flag.StringVar(&filename, "f", "", "filename to output the record (if not specified, output to stdout)")
	flag.Parse()

	if filename == "" {
		loggen.GenerateLog(days, errRate)
	} else {
		loggen.GenerateLogToFile(days, errRate, filename)
	}
}
