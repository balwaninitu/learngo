package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "" + "lazy cat jumps over and over again"

func main() {

	words := strings.Fields(corpus)

	query := os.Args[1:]

queries:

	for _, q := range query {

		q = strings.ToLower(q)

		for i, w := range words {

			if w == q {

				fmt.Printf("#%-2d: %q\n", i+1, w)

				break queries
			}

		}

	}

}
