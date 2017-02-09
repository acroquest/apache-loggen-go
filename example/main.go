package main

import (
	"flag"
	"github.com/acroquest/apache-loggen-go"
)

func main() {
	var (
		days    int
		errRate float64
	)
	flag.IntVar(&days, "day", 1, "days to output")
	flag.IntVar(&days, "d", 1, "same as -day")
	flag.Float64Var(&errRate, "err", 0.1, "error rate")
	flag.Float64Var(&errRate, "e", 0.1, "same as -err")
	flag.Parse()

	loggen.GenerateLog(days, errRate)
}
