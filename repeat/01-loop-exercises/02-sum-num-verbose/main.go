package main

import "fmt"

func main() {

	var sum int

	min, max := 1, 10

	for i := min; i <= max; i++ {

		sum += i

		fmt.Print(i)

		if i < max {

			fmt.Print(" + ")
		}

	}

	fmt.Printf(" = %d", sum)

}
