package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 || len(args) > 5 {

		fmt.Println("give me 5 no")

		return

	}

	var nums [5]float64

	for i, v := range args {

		n, err := strconv.ParseFloat(v, 64)

		if err != nil {

			continue
		}

		nums[i] = n
	}

	for range nums {

		for i, v := range nums {

			if i < len(nums)-1 && v > nums[i+1] {

				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}

	}

	fmt.Println(nums)

}
