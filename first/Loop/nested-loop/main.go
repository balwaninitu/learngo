package main

import "fmt"

const max = 10

func main() {

	fmt.Printf("%10s", "X")

	for i := 0; i <= max; i++ {

		fmt.Printf("%10d", i)
	}

	fmt.Println()

	for i := 0; i <= max; i++ {

		fmt.Printf("%10d", i)

		for j := 0; j <= max; j++ {

			fmt.Printf("%10d", i*j)
		}

		fmt.Println()
	}

}
