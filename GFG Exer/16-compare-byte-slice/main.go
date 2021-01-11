package main

import (
	"bytes"
	"fmt"
)

func main() {

	slc1 := []byte{'g', 'o', 's', 'c'}

	slc2 := []byte{'g', 'o'}

	s := bytes.Compare(slc1, slc2)

	if s == 0 {

		fmt.Println("they are equal")
	} else {

		fmt.Println("not equal")
	}

}
