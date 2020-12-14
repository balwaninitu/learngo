package main

import "fmt"

func main() {

	var (
		perimeter int

		length, width = 6, 5
	)

	perimeter = 2 * (length + width)

	fmt.Println(perimeter)

}
