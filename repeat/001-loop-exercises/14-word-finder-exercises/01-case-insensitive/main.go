package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	const quote = "lets have party tonight and over the moon and over high"

	words := strings.Fields(quote)

	args := os.Args[1:]

	for _, a := range args {

		for i, w := range words {

			if a == w {

				fmt.Printf("%-2d: %q\n", i+1, w)

				break

			}

		}

	}

}
