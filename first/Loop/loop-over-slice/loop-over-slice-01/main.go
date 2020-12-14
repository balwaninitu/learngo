package main

import (
	"fmt"
	"strings"
)

func main() {

	words := strings.Fields(
		"Dont count your chickens before they are hatched")

	for j := 0; j < len(words); j++ {

		fmt.Printf("#%-2d: %q\n", j+1, words[j])
	}

}
