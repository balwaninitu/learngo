package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		names    [3]string
		distance [5]int
		data     [5]byte
		ratios   [1]float64
		alives   [4]bool
		zero     [0]byte
	)

	names[0] = "Lina"
	names[1] = "Mina"
	names[2] = "Tina"

	distance[0] = 18
	distance[1] = 25
	distance[2] = 58
	distance[3] = 78
	distance[4] = 41

	data[0] = 'H'
	data[1] = 'E'
	data[2] = 'L'
	data[3] = 'L'
	data[4] = 'O'

	ratios[0] = 3.14

	alives[0] = true
	alives[01] = true
	alives[2] = false
	alives[3] = false

	_ = zero

	separator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("names", separator)

	for i := 0; i < len(names); i++ {

		fmt.Printf("names[%d]:  %q\n", i, names[i])

	}

	fmt.Print("\ndistamce", separator)

	for i := 0; i < len(distance); i++ {

		fmt.Printf("distance[%d]: %d\n", i, distance[i])
	}

	fmt.Print("\ndata", separator)

	for i := 0; i < len(data); i++ {

		fmt.Printf("data[%d]: %d\n", i, data[i])
	}

	fmt.Print("\nratios", separator)

	for i := 0; i < len(ratios); i++ {

		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}

	fmt.Print("\nalives", separator)

	for i := 0; i < len(alives); i++ {

		fmt.Printf("\nalives[%d]: %t\n", i, alives[i])
	}

	fmt.Print("\nzero", separator)

	for i := 0; i < len(zero); i++ {

		fmt.Printf("/nzero[%d]: %d\n", i, zero[i])
	}

	fmt.Print("names", separator)

	for i, v := range names {

		fmt.Printf("names[%d]: %s\n", i, v)
	}

	fmt.Print("\ndistamce", separator)

	for i, v := range distance {

		fmt.Printf("\ndistance [%d]: %d\n", i, v)
	}

	fmt.Print("\ndata", separator)

	for i, v := range data {

		fmt.Printf("\ndata[%d]: %d\n", i, v)
	}

	fmt.Print("\nratios", separator)

	for i, v := range ratios {

		fmt.Printf("ratios[%d]: %.2f\n", i, v)
	}

	fmt.Print("\nalives", separator)

	for i, v := range alives {

		fmt.Printf("\nalives[%d]: %t\n", i, v)
	}

	fmt.Print("\nzero", separator)

	for i, v := range zero {

		fmt.Printf("\n[%d]: %d\n", i, v)
	}

}
