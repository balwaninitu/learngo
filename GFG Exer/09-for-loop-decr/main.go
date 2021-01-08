package main

import "fmt"

func main() {

	var i int

	for i = 10; i > 1; i-- {

		i -= 3
	}

	fmt.Println("i =", i)

}
