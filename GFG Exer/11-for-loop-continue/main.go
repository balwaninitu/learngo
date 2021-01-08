package main

import "fmt"

func main() {

	var i int

	for i = 0; i < 15; i++ {

		if i == 7 {

			i = 3 + i

			continue
		}

		fmt.Println("i", i)
	}

}
