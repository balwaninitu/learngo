package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// fmt.Println(a)

	words := filepath.SplitList(os.Getenv("PATH"))

	// fmt.Println(words)

	args := os.Args[1:]

	for _, q := range args {

		for i, w := range words {

			if strings.Contains(w, q) {

				fmt.Printf("%-2d: %q\n", i+1, w)

				break
			}
		}
	}

}
