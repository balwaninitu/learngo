package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {

		fmt.Println("Give me a letter")

		return
	}

	s := args[1]

	if strings.IndexAny(s, "aeiou") != -1 {

		fmt.Printf("%q is a vowel", s)
	} else if s == "y" || s == "w" {

		fmt.Printf("%q is sometimes a vowel, sometimes not\n", s)

	} else {

		fmt.Printf("%q is consonant", s)
	}

}
