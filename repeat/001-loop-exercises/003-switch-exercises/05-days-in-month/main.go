package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	if len(os.Args) != 2 {

		fmt.Println("Give me a month name")

		return
	}

	days, m := 28, os.Args[1]

	year := time.Now().Year()

	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	switch m {

	case "jan", "mar", "may", "jul", "aug", "oct", "dec":

		days = 31

	case "apr", "jun", "sep", "nov":

		days = 30

	case "feb":

		if leap {
			days = 29
		}

	default:

		fmt.Printf("%s is not a month", m)

		return

	}

	fmt.Printf("%s has %d days\n", m, days)

}
