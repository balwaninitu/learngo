package main

import "fmt"

func main() {

	data := [5]uint8{

		2 * 4, 1 * 6, 5 * 0, 3 * 6, 8 * 1,
	}

	fmt.Printf("%#v\n", data)

	fmt.Printf("%T\n", data)

	fmt.Printf("%d\n", data)

}
