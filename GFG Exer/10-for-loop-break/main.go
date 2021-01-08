package main

import "fmt"

func main() {

	var i int

	for i = 0; i <= 5; i++ {

		if i == 3 {

			break
		}
	}

	fmt.Println("i =", i)

}
