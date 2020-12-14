package main

import (
	"fmt"
	"os"
)

func main() {

	length := len(os.Args[1])

	fmt.Println(length)

}
