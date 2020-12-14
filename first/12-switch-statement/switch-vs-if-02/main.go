package main

import (
	"fmt"
	"os"
)

func main() {

	switch m := os.Args[1]; m {

	case "dec", "jan", "feb":

		fmt.Println("winter")

	case "mar", "apr", "may":

		fmt.Println("spring")

	case "jun", "jul", "aug":

		fmt.Println("summer")

	case "sept", "oct", "nov":

		fmt.Println("fall")

	default:

		fmt.Printf("%q is not a month name.\n", m)

	}

}
