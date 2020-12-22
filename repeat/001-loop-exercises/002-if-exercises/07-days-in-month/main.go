package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {

		fmt.Println("Give me a month name")

		return
	}

	m := strings.ToLower(os.Args[1])

	leap := "feb has 29 days"

	if m == "feb" && m == leap {

		fmt.Printf("%s has 29 days ", m)

	} else if m == "feb" {

		fmt.Printf("%s has 28 days ", m)
	} else if m == "jan" || m == "mar" || m == "may" || m == "jul" || m == "aug" || m == "oct" || m == "dec" {

		fmt.Printf("%s has 31 days ", m)

	} else if m == "apr" || m == "jun" || m == "sep" || m == "nov" {

		fmt.Printf("%s has 30 days ", m)

	} else {

		fmt.Printf("%s is not a month ", m)
	}

}
