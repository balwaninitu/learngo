package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	data := "2 4 6 1 3 5"

	splitted := strings.Fields(data)

	fmt.Println(splitted)

	var nums []int

	for _, s := range splitted {

		n, _ := strconv.Atoi(s)

		nums = append(nums, n)
	}

	fmt.Println("nums:", nums)

	even, odd := nums[:3], nums[3:]

	fmt.Println("Even", even)

	fmt.Println("Odd", odd)

	twoNums := nums[2:4]

	fmt.Println("twoNums", twoNums)

	firstTwo := nums[:2]

	fmt.Println("firstTwo", firstTwo)

	lastTwo := nums[(len(nums))-2:]

	fmt.Println("lastTwo:", lastTwo)

	fmt.Println("last even one:", even[(len(even))-1])

	fmt.Println("last odd two:", odd[(len(odd))-2:])

}
