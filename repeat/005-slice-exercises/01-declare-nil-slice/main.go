package main

import "fmt"

func main() {

	var (
		names []string

		distance []int

		data []byte

		ratio []float64

		alives []bool
	)

	fmt.Printf("Type :%T length :%d equal :%t\n", names, len(names), names == nil)

	fmt.Printf("Type :%T length :%d equal :%t\n", distance, len(distance), distance == nil)

	fmt.Printf("Type :%T length :%d equal :%t\n", data, len(data), data == nil)

	fmt.Printf("Type :%T length :%d equal :%t\n", ratio, len(ratio), ratio == nil)

	fmt.Printf("Type :%T length :%d equal :%t\n", alives, len(alives), alives == nil)

}
