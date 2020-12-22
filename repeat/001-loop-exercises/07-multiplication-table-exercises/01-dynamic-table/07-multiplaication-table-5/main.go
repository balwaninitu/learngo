package main

import "fmt"

func main() {

	fmt.Printf("%5s", "X")

	for i := 0; i <= 5; i++ {

		fmt.Printf("%5d", i)

	}

	fmt.Println()

	for i := 0; i <= 5; i++ {

		fmt.Printf("%5d", i)

		for j := 0; j <= 5; j++ {

			fmt.Printf("%5d", j*i)
		}

		fmt.Println()
	}

}
