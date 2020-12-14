package main

import (
	"fmt"

	"os"
)

func main() {

	args := os.Args

	if len(args) != 3 {

		fmt.Println("usage : [username]  [password]")

		return
	}

	u := args[1]
	p := args[2]

	if u != "jack" {

		fmt.Printf("Acess denied for %q.\n", u)

	} else if p != "1888" {

		fmt.Printf("invalid password for %q.\n", u)

	} else {

		fmt.Printf("Access granted for %q.\n", u)
	}

}
