package main

import "fmt"

func main() {

	var (
		phones map[string]string

		product map[int]bool

		multiphone map[string][]string

		basket map[int]map[int]int
	)

	fmt.Printf("%#v\n", phones)

	fmt.Printf("%#v\n", product)

	fmt.Printf("%#v\n", multiphone)

	fmt.Printf("%#v\n", basket)

}
