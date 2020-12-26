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

	names = []string{"ina", "mina", "dica"}

	distance = []int{5, 4, 3, 2, 7}

	data = []byte{'h', 'e', 'l', 'l', 'o'}

	ratio = []float64{5.4, 7.0, 8.1}

	alives = []bool{true, false}

	fmt.Printf("Type :%T length :%d equal :%t\n", names, len(names), names == nil)

	fmt.Printf("Type :%T length :%d equal :%t\n", distance, len(distance), distance == nil)

	fmt.Printf("Type :%T length :%d equal :%t\n", data, len(data), data == nil)

	fmt.Printf("Type :%T length :%d equal :%t\n", ratio, len(ratio), ratio == nil)

	fmt.Printf("Type :%T length :%d equal :%t\n", alives, len(alives), alives == nil)

}
