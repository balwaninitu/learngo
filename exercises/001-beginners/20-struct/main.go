package main

import "fmt"

func main() {

	type condo struct {
		name string

		blks int

		new bool

		sqft float64
	}

	first := condo{"Aquarius", 10, false, 6.8}

	second := condo{"Clearwater", 12, false, 5.5}

	third := condo{"WaterGold", 10, true, 4.2}

	fmt.Println(condo(first))

	fmt.Println(condo(second))

	fmt.Println(condo(third))

	second.blks = 14

	fmt.Println(condo(second))

	first.new = true

	fmt.Println(condo(first))

	p := &condo{}

	fmt.Println(p)

}
