package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {

		fmt.Println("Tell me the magnitude of the earthquake in human terms ")

		return
	}

	r := os.Args[1]

	var res string

	switch r {

	case "micro":

		res = "less than 2.0"
	case "very minor":

		res = "2- 2.9"
	case "minor":

		res = "3- 3.9"
	case "light":

		res = "4- 4.9"
	case "moderator":

		res = "5- 5.9"
	case "strong":

		res = "6- 6.9"
	case "major":

		res = "7- 7.9"
	case "great":

		res = "8- 8.9"
	case "massive":

		res = "10+"
	default:

		res = "unknown"

	}

	fmt.Printf("%s's richter scale is %s\n", r, res)

}
