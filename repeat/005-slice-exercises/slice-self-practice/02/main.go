package main

import "fmt"

func main() {

	// print value & index

	num := []int{5, 3, 2, 8, 2}

	var sum int

	for i, v := range num {

		num[i] = i

		sum += v

		fmt.Printf("num[%d]: %d\n", num[i], v)

		fmt.Println()

	}

	fmt.Printf("sum = %d average %d\n", sum, sum/len(num))

}
