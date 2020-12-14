package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {

		fmt.Println("Tell me a book title")

		return
	}

	Books := [4]string{

		"Kafka's revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	title := strings.ToLower(args[0])

	fmt.Println("Search Result:")

	var found bool

	for _, v := range Books {

		if strings.Contains(strings.ToLower(v), title) {

			fmt.Println("+", v)

			found = true
		}
	}

	if !found {

		fmt.Printf("We dont have the book: %q\n", title)
	}

}
