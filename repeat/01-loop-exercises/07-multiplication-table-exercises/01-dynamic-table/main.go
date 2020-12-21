package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args

	if len(args) != 2 {

		fmt.Println("Give me the size of table")

		return
	}

	n, err := strconv.Atoi(args[1])

	if err != nil || n < 0 {

		fmt.Println("Wrong size")

		return
	}

	fmt.Printf("%5s", "X")

	for i := 0; i <= n; i++ {

		fmt.Printf("%5d", i)

	}

	fmt.Println()

	for i := 0; i <= n; i++ {

		fmt.Printf("%5d", i)

		for j := 0; j <= n; j++ {

			fmt.Printf("%5d", i*j)
		}

		fmt.Println()
	}

}
