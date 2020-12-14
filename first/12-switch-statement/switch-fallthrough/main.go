package main

import "fmt"

func main() {

	i := 123

	switch {

	case i > 100:

		fmt.Println("Big")

		fallthrough

	case i < 100:

		fmt.Println("Positive")

		fallthrough

	default:

		fmt.Println("number")

	}

}
