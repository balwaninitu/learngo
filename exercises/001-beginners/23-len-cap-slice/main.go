package main

import "fmt"

func main() {

	s := []int{5, 90, 45, 23, 12, 78, 56, 6, 3, 2}

	s = s[:6]

	printSlice(s)

	s = s[1:5]

	printSlice(s)

	s = s[2:4]

	printSlice(s)

	s = s[1:]

	printSlice(s)

}

func printSlice(s []int) {

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
