package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	data := "2 4 6 1 3 5"

	split := strings.Fields(data)

	var nums []int

	for _, v := range split {

		n, _ := strconv.Atoi(v)

		nums = append(nums, n)
	}

	fmt.Println("nums     :", nums)

	evens, odd := nums[:3], nums[3:]

	fmt.Println("Even: ", evens)

	fmt.Println("Even: ", odd)

	midle := nums[2:4]

	first := nums[:2]

	l := len(nums)

	last := nums[l-2:]

	evelast := evens[len(evens)-1:]

	oddlast := odd[len(odd)-2:]

	fmt.Println("Middle: ", midle)

	fmt.Println("first: ", first)

	fmt.Println("last: ", last)

	fmt.Println("evelast: ", evelast)

	fmt.Println("oddlast: ", oddlast)

}
