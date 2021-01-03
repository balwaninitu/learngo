package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {

		fmt.Println("provide a number")

		return

	}

	n, err := strconv.Atoi(os.Args[1])

	if err != nil || n < 0 {

		fmt.Println("Err")

	} else if n <= 10 {

		fmt.Printf("%d number is between 1 and 10\n", n)
	} else {

		fmt.Println("exceed limit")
	}

}
