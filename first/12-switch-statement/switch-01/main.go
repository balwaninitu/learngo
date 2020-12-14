package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	i, _ := strconv.Atoi(os.Args[1])

	switch {
	case i > 0:

		fmt.Println("Number is positive")

	case i < 0:

		fmt.Println("Number is negative")

	default:
		fmt.Println("Number is zero")

	}

}
