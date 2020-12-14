package main

import (
	"fmt"
	"math/rand"
)

func main() {

	rand.Seed(100)

	guess := 100

	for n := 0; n != guess; {

		n = rand.Intn(guess + 1)

		fmt.Printf("%d ", n)
	}

	fmt.Println()

}
