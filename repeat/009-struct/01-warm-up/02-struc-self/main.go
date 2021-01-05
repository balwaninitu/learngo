package main

import "fmt"

func main() {

	type person struct {
		name       string
		age        int
		occupation string
		height     float64
		male       bool
	}

	p := person{"nitu", 37, "coder", 5.1, false}

	fmt.Println(p.age)

	fmt.Println(p.male)
}
