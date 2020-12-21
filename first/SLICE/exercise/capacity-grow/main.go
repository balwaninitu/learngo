package main

import "fmt"

func main() {

	new := []int{}

	fmt.Printf("len:%d cap: %d", len(new), cap(new))

	for len(new) < 10e6 {

		c := cap(new)

		new = append(new, 1)

		fmt.Printf("len: %d\n cap:%d", len(new), c)
	}

}
