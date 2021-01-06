package main

import "fmt"

func main() {

	nums := [...]int{5, 78, 3, 9, 12, 5}

	fmt.Println(nums)

	a := nums[:3]

	b := nums[1:]

	x := nums[1]

	fmt.Println("x", x)

	fmt.Println(a, b)

	nums[4] = 23

	fmt.Println(nums)

	fmt.Println(a, b)

	st := []struct {
		s string
		i int
	}{
		{"Tia", 24},
		{"Pia", 23},
		{"Nia", 25},
		{"kia", 23},
	}

	fmt.Println(st)

}
