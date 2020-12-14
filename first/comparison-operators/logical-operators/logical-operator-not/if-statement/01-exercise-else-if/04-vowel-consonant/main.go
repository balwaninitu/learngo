package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {

		fmt.Println("Give me a letter")
		return
	}

	s := args[1]

	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" {

		fmt.Printf("%q is a vowel.\n", s)

	} else if s == "y" || s == "w" {

		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)

	} else {

		fmt.Printf("%q is a consonant.\n", s)
	}
}
