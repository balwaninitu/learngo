package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args

	if len(args) != 2 {

		fmt.Println("Pick a number")

		return
	}

	n, err := strconv.Atoi(args[1])

	if err != nil || n < 0 {

		fmt.Println("its not valid number")

		return
	}

	if n%8 == 0 {

		fmt.Printf("%d is an even number and its divisible by 8", n)

	} else if n%2 == 0 {

		fmt.Printf("%d is an even number ", n)

	} else {

		fmt.Printf("%d is an odd number", n)
	}

}
