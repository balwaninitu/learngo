package main

import (
	"fmt"
	"os"
)

func main() {

	city := os.Args[1]

	switch city {
	case "Paris", "Lyon":

		fmt.Println("France")

	case "Tokyo", "Kyoto":

		fmt.Println("Japan")

	default:

		fmt.Println("Where?")

	}

}
