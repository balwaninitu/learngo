package main

import "fmt"

func sum(x, y int) int {

	return x + y

}

func main() {

	x, y := 5, 6

	a := sum(x, y)

	fmt.Println(a)

}
