package main

import (
	"flag"
)

func main() {
	var day int
	flag.IntVar(&day, "day", 1, "day on advent of code")
	flag.Parse()

	switch day {
	case 1:
		day1()
	case 2:
		day2()
	}
}
