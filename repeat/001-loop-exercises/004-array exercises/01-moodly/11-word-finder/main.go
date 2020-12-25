package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

	args := os.Args[1:]

	if len(args) == 0 {

		fmt.Println("Please give me a word to search.")

		return
	}

	filter := [...]string{"and", "or", "was", "the", "since", "very"}

	words := strings.Fields(strings.ToLower(corpus))

oLoop:
	for _, a := range args {

		a = strings.ToLower(a)

		for _, v := range filter {

			if a == v {

				continue oLoop
			}

			for i, w := range words {

				if w == a {

					fmt.Printf("%-2d : %q\n", i+1, w)

					break
				}
			}
		}
	}

}
