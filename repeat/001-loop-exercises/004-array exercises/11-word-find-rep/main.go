package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

	args := os.Args[1:]

	if len(args) != 1 {

		fmt.Println("Please give me a word to search.")

		return
	}

	words := strings.Fields(strings.ToLower(corpus))

	filters := [...]string{"and", "or", "was", "the", "since", "very"}

oLoop:

	for _, q := range args {

		q = strings.ToLower(q)

		for _, v := range filters {

			if q == v {

				continue oLoop
			}

		}

		for i, w := range words {

			if q == w {

				fmt.Printf("%d: %q\n", i, w)

			}

		}

	}

}
