package main

import (
	"fmt"
	"strings"

	"os"
)

func main() {

	msg := os.Args[1]

	s := msg + strings.Repeat("!", len(msg))

	fmt.Println(s)

}
