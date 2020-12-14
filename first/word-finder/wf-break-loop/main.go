package main

import (
	"fmt"
	"os"
	"strings"
)

const corpse = "" + "lazy cat jumps over and over again"

func main() {

	words := strings.Fields(corpse)

	query := os.Args[1:]

	for _, q := range query {

		q = strings.ToLower(q)

	search:

		for i, w := range words {

			switch q {

			case "and", "or", "the":

				break search
			}

			if q == w {

				fmt.Printf("#%-2d: %q\n", i+1, w)

			}

		}

	}

}
