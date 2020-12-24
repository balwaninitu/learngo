package main

import "fmt"

func main() {

	names := [...]string{"ina", "mina", "dica"}

	distance := [5]int{4, 5, 3, 2, 1}

	data := [5]byte{'h', 'e', 'l', 'l', 'o'}

	ratios := [1]float64{2.5}

	alives := [4]bool{true, false, true, true}

	zero := [0]byte{}

	for i := 0; i < len(names); i++ {

		fmt.Printf("names %d : %q\n", i, names[i])

	}

	fmt.Println()

	for i := 0; i < len(distance); i++ {

		fmt.Printf("distance %d : %d\n", i, distance[i])

	}

	fmt.Println()

	for i := 0; i < len(data); i++ {

		fmt.Printf("%d data: %q\n", i, data[i])
	}

	fmt.Println()

	for i := 0; i < len(ratios); i++ {

		fmt.Printf("%d ratios : %f\n", i, ratios[i])
	}

	fmt.Println()

	for i := 0; i < len(alives); i++ {

		fmt.Printf("%d alives : %t\n", i, alives[i])
	}

	fmt.Println()

	for i := 0; i < len(zero); i++ {

		fmt.Printf("%d zero : %q\n", i, zero[i])
	}

	fmt.Println()

	for i, n := range names {

		fmt.Printf("%d names : %q\n", i, n)
	}

	fmt.Println()

	for i, d := range distance {

		fmt.Printf("[%d] distance: %d\n", i, d)
	}

	fmt.Println()

	for i, da := range data {

		fmt.Printf("[%d] data : %q\n", i, da)
	}

	fmt.Println()

	for i, r := range ratios {

		fmt.Printf("[%d] ratios : %f\n", i, r)
	}

	fmt.Println()

	for i, a := range alives {

		fmt.Printf("[%d] alives : %t\n", i, a)
	}

	fmt.Println()

	for i, z := range zero {

		fmt.Printf("[%d] zero : %q\n", i, z)
	}

	fmt.Println()

}
