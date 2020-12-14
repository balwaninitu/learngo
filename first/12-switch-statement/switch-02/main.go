package main

import "fmt"

func main() {

	i := 0

	switch {

	case i > 100:

		fmt.Println("Big positive number")

	case i < 0:

		fmt.Println("positive number")

	default:

		fmt.Println("number")
	}

}
