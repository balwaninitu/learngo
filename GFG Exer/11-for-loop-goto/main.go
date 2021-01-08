package main

import "fmt"

func main() {

	i := 0

loop:
	for i < 15 {

		if i == 5 {

			i = i + 3

			goto loop
		}

		fmt.Println("i =", i)

		i++
	}

}
