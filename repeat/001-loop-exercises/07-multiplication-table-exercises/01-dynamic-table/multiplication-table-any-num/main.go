package main

import "fmt"

func main() {

	fmt.Printf("%5s", "X")

	var num int

	num = -9

	for i := 0; i <= num; i++ {

		fmt.Printf("%5d", i)

	}

	fmt.Println()

	for i := 0; i <= num; i++ {

		fmt.Printf("%5d", i)

		for j := 0; j <= num; j++ {

			fmt.Printf("%5d", j)
		}

		fmt.Println()

	}

}
