package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {

		fmt.Println("Send me some items and I will sort them")

		return
	}

	sort.Strings(args)

	fmt.Println(args)

	c = ':)'

}
