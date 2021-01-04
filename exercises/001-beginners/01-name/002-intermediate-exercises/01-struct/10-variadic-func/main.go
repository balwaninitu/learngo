package main

import "fmt"

func sum(nums ...int) {

	total := 0

	for _, num := range nums {

		total += num
	}

	fmt.Println("total", total)

}

func main() {

	sum(5, 7, 2, 8, 4, 1)

}
