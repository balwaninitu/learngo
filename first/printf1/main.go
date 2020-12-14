package main

import "fmt"

func main() {

	var (
		name   = "Riddham"
		age    = 8
		class  = 3
		clever = true
		height = 120
		weight = 19.5
	)

	fmt.Printf("My daughter name is %[1]v\nShe is %[2]v years old.\n", name, age)
	fmt.Printf("She will go to %v next year, is she clever? Yes %v.\n", class, clever)
	fmt.Printf("Her height is %vcm and weight is %vkg.\n", height, weight)
	fmt.Printf("She is apple of my eyes.")
}
