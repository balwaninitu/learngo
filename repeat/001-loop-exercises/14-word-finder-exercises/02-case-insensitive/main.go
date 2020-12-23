package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	const quote = `jingle bell jingle bell jingle all the way 
oh what fun it is to ride in the one horse open sleigh`

	words := strings.Fields(quote)

	args := os.Args[1:]
oLoop:
	for _, a := range args {

		a = strings.ToLower(a)
	iLoop:
		for i, w := range words {

			if a == "and" || a == "or" || a == "the" {

				break iLoop
			}

			if a == w {

				fmt.Printf("%-2d : %q\n", i+1, w)

				continue oLoop
			}
		}
	}

}
