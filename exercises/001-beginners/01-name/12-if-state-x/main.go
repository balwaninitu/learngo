package main

import "fmt"

func main() {

	x := 5

	if x > 2 {

		x = x / 2
	}

	if x < 2 {

		x = x / 1
	}

	fmt.Println(x)
}
