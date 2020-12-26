package main

import "fmt"

func main() {
	toppings := []string{"black olives", "green peppers"}

	var pizza []string
	pizza = append(pizza, toppings...)
	pizza = append(toppings, "onions", "extra cheese")

	fmt.Printf("toppings:       : %s\n", pizza)
}
