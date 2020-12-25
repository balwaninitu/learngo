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

	w := strings.ToLower(args[0])

	for _, v := range titles {

		if strings.Contains(strings.ToLower(v), w) {

			fmt.Println("+", v)

			return

		}

	}
	fmt.Printf("dont have %q\n", w)
}
