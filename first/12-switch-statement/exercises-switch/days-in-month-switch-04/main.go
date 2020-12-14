package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	args := os.Args

	if len(args) != 2 {

		fmt.Println("Give me a month name")

		return
	}

	days, m := 28, args[1]

	year := time.Now().Year()

	leap := year%4 == 0 && (year%100 != 100 || year%400 == 0)

	switch strings.ToLower(m) {

	case "jan", "mar", "may", "jul", "aug", "oct", "dec":

		days = 31

	case "apr", "jun", "sept", "nov":

		days = 30

	case "feb":

		if leap {

			days = 29
		}

	default:

		fmt.Printf("%q is not a month.\n", m)

	}

	fmt.Printf("%q has %d days.\n", m, days)

}
