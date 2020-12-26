package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	titles := [...]string{

		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	args := os.Args[1:]

	if len(args) != 1 {

		fmt.Println("Tell me a book title")

		return
	}

	s := strings.ToLower(args[0])

	for _, v := range titles {

		if strings.Contains(strings.ToLower(v), s) {

			fmt.Println("+", v)

			return

		}
	}

	fmt.Println("We don't have the book:", s)

}
