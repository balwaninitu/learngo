package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args

	if len(args) != 2 {

		fmt.Println("pick a number")

		return
	}

	n, err := strconv.Atoi(args[1])

	if err != nil {

		fmt.Printf("%q is not a number", args[1])

		return

	}

	if n%8 == 0 {

		fmt.Printf("%d is an even number and its divisible by 8.\n", n)

	} else if n%2 == 0 {

		fmt.Printf("%d is an even number", n)

	} else {

		fmt.Printf("%d is an odd number", n)
	}

}
