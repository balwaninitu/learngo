package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]

	l := len(args)

	if l != 5 {

		fmt.Println("Please tell me numbers (maximum 5 numbers).")

		return

	}

	var (
		num [5]float64

		sum float64

		total float64
	)

	for i, v := range args {

		n, err := strconv.ParseFloat(v, 64)

		if err != nil {
			continue
		}

		total++
		sum += n

		num[i] = n

	}

	fmt.Println("your number", num)

	fmt.Println("average :", sum/total)

}
