package main

import "fmt"

func main() {

	var (
		names []string

		distance []int

		data []byte

		ratios []float64

		alives []bool
	)

	names = []string{"lina", "mina", "tina"}

	distance = []int{4, 5, 6, 2, 1}

	data = []byte{'h', 'e', 'y'}

	ratios = []float64{67.7, 8.9, 9.0, 5.7}

	alives = []bool{true, false, false}

	fmt.Printf("names: %T %d %t\n", names, len(names), names == nil)

	fmt.Printf("distance: %T %d %t\n", distance, len(distance), distance == nil)

	fmt.Printf("data: %T %d %t\n", data, len(data), data == nil)

	fmt.Printf("ratios: %T %d %t\n", ratios, len(ratios), ratios == nil)

	fmt.Printf("alives: %T %d %t\n", alives, len(alives), alives == nil)

}
