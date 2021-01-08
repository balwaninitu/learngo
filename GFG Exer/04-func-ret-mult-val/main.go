package main

import "fmt"

func name(x, y int) (rect int, sq int) {

	rect = x * y

	sq = x * x

	return
}

func main() {

	var s1, s2 = name(3, 7)

	fmt.Println("rect", s1)

	fmt.Println("sq", s2)

}
