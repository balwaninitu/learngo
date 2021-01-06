package main

import "fmt"

func main() {

	a := make([]int, 5)

	printSlice("a", a)

	b := make([]int, 0, 5)

	printSlice("b", b)

	c := b[1:3]

	printSlice("c", c)

	d := c[:]

	printSlice("d", d)

}

func printSlice(s string, i []int) {

	fmt.Printf("%s len=%d cap=%d %v\n", s, len(i), cap(i), i)

}
