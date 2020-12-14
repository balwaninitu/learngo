package main

import "fmt"

func main() {

	var i int

loop:

	if i < 3 {

		fmt.Println("Looping")

		i++

		goto loop
	}

	fmt.Println("done")

}
