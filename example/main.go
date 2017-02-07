package main

import (
	"flag"
	"github.com/acroquest/apache-loggen"
)

func main() {
	var (
		days    int
		errRate int
	)
	flag.IntVar(&days, "day", 1, "days to output")
	flag.IntVar(&days, "d", 1, "same as -day")
	flag.IntVar(&errRate, "err", 1, "error rate")
	flag.IntVar(&errRate, "e", 1, "same as -err")
	flag.Parse()

	loggen.GenerateLog(days)
}
