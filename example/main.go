package main

import (
	"flag"
	"github.com/acroquest/apache-loggen-go"
	"strconv"
	"strings"
)

func main() {
	var (
		days       int
		errRate    float64
		filename   string
		numofFiles int
	)

	flag.IntVar(&days, "d", 1, "same as -day")
	flag.Float64Var(&errRate, "e", 0.1, "same as -err")
	flag.StringVar(&filename, "f", "", "filename to output the record (if not specified, output to stdout)")
	flag.IntVar(&numofFiles, "n", 1, "number of files to output")
	flag.Parse()

	if filename == "" {
		loggen.GenerateLog(days, errRate)
	} else {
		splitted := strings.Split(filename, ".")
		if len(splitted) == 1 {
			for i := 1; i <= numofFiles; i++ {
				f := filename + "-" + strconv.Itoa(i)
				loggen.GenerateLogToFile(days, errRate, f)
			}
		} else {
			head := strings.Join(splitted[:len(splitted)-1], "")
			for i := 1; i <= numofFiles; i++ {
				f := head + "-" + strconv.Itoa(i) + "." + splitted[len(splitted)-1]
				loggen.GenerateLogToFile(days, errRate, f)
			}
		}
	}
}
