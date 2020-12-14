package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	if len(args) != 2 {

		fmt.Println("Tell me the magnitude of the Earthquake in human terms")

		return
	}

	r := args[1]

	var richter string

	switch r {

	case "micro":

		richter = "less than 2.0"

	case "very minor":

		richter = "2 - 2.9"

	case "minor":

		richter = " 3 - 3.9"

	case "light":

		richter = " 4 - 4.9"

	case "moderate":

		richter = " 5 - 5.9"

	case "strong":

		richter = " 6 - 6.9"

	case "major":

		richter = "  7 - 7.9"

	case "great":

		richter = " 8 - 8.9"

	case "massive":

		richter = " 10+"

	default:

		richter = "unknown"
	}

	fmt.Printf("%s's richter scale is %s\n", r, richter)

}