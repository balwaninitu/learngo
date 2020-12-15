package main

import (
	"fmt"
	"time"
)

func main() {

	var pizza []string

	pizza = append(pizza, "cheese", "olives", "pineapple")

	var departure []time.Time

	now := time.Now()
	departure = append(departure, now,
		now.Add(time.Hour*24),
		now.Add(time.Hour*48))

	var years []int

	years = append(years, 2010, 2017, 2005)

	var lights []bool

	lights = append(lights, true, false, true)

	fmt.Printf("Pizza :      %s\n", pizza)

	fmt.Printf("departure :      %s\n", departure)

	fmt.Printf(" years :      %d\n", years)

	fmt.Printf("lights :      %t\n", lights)

}
