package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	if len(os.Args) != 2 {

		fmt.Println("Give me a month name")

		return
	}

	year := time.Now().Year()

	leap := year%4 == 0 && (year%100 != 0 || year%400 == 00)

	days := 28

	month := os.Args[1]

	m := strings.ToLower(month)

	if m == "april" || m == "june" || m == "september" || m == "november" {

		days = 30

		fmt.Printf("%q has %d days.\n", month, days)

	} else if m == "january" || m == "march" || m == "may" || m == "july" || m == "august" || m == "october" || m == "december" {

		days = 31

		fmt.Printf("%q has %d days.\n", month, days)

	} else if m == "february" {

		if leap {

			days = 29
		}

		fmt.Printf("%q has %d days.\n", month, days)

	} else {

		fmt.Printf("%q is not a month.\n", month)
	}
}
