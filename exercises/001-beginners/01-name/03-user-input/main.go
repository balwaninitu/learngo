package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 {

		fmt.Println("Enter your name")

	} else {

		fmt.Printf("name: %q\n", os.Args[1:])
	}

}
