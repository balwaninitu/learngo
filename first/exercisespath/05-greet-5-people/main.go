package main

import (
	"fmt"

	"os"
)

func main() {

	var (
		a = len(os.Args) - 1
		b = os.Args[1]
		c = os.Args[2]
		d = os.Args[3]
	)

	fmt.Println("there are", a, "people !")
	fmt.Println("Hello", b, "!")
	fmt.Println("Hello", c, "!")
	fmt.Println("Hello", d, "!")

}
