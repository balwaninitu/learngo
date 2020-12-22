package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {

		fmt.Println("Give me args")
	}

	l := len(args)

	if l == 1 {

		fmt.Printf("There is one : %q\n", args[0])

	} else if l == 2 {

		fmt.Printf("There are two : %q %q\n", args[0], args[1])

	} else {

		fmt.Printf("There are %d arguments\n", l)
	}

}
