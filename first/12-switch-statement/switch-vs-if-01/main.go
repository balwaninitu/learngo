package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	if len(args) != 2 {

		fmt.Println("Give me a month name")

		return
	}

	m := args[1]

	if m == "Dec" || m == "Jan" || m == "Feb" {

		fmt.Println("Winter")
	} else if m == "Mar" || m == "Apr" || m == "May" {

		fmt.Println("Spring")

	} else if m == "Jun" || m == "Jul" || m == "Aug" {

		fmt.Println("Summer")

	} else if m == "Sept" || m == "Oct" || m == "Nov" {

		fmt.Println("Fall")

	} else {

		fmt.Printf("%q is not a month", m)
	}

}
