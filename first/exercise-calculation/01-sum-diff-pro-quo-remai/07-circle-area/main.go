package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		radius = 10.
		area   float64
	)
	area = math.Pi * math.Pow(radius, 2)

	fmt.Printf("radius: %g -> area : , %g\n", radius, area)

	fmt.Printf("radius: %g -> area : , %.2f\n", radius, area)

}
