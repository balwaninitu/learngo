package main

import "fmt"

func main() {

	//iterate loop for slice elements

	slice1 := []int{4, 3, 2, 1, 5}

	for i := 0; i <= len(slice1); i++ {

		fmt.Printf("slice[%d]: %d\n", i, slice1[i])

	}

}
