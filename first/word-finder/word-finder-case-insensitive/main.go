package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "" + "mar had a little lamb, little lamb, little lamb its fleece was white as snow "

func main() {

	words := strings.Fields(corpus)

	query := os.Args[1:]

	for _, q := range query {

		q = strings.ToLower(q)

		for i, w := range words {

			if q == w {

				fmt.Printf("#%-2d : %q\n", i+1, w)

				break
			}

		}
	}
}
