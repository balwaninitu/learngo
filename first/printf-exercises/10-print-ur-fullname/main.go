package main

import (
	"fmt"
	"os"
)

func main() {

	name, lastname := os.Args[1], os.Args[2]

	msg := "My name is %s and my last name is %s.\n"

	fmt.Printf(msg, name, lastname)

}
