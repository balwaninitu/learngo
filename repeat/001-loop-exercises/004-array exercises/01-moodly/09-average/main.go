package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]

	if l := len(args); l == 0 || l > 5 {

		fmt.Println(" Please tell me numbers (maximum 5 numbers)")

		return
	}

	var (
		sum float64

		total float64

		nums [5]float64
	)

	for i, v := range args {

		a, err := strconv.ParseFloat(v, 64)

		if err != nil {

			continue
		}

		total++
		nums[i] = a

		sum += a

	}

	fmt.Println("your numbers :", nums)

	fmt.Println("average :", sum/total)

}
