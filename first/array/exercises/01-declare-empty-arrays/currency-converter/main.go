package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	const (
		EUR = iota
		SGD
		INR
	)

	rates := [...]float64{

		EUR: 5.90,
		SGD: 20.02,
		INR: 60.09,
	}

	args := os.Args[1:]

	if len(args) != 1 {

		fmt.Println("Please provide amount to be converted")

		return
	}

	amount, err := strconv.ParseFloat(args[0], 64)

	if err != nil {

		fmt.Println("Invalid")

		return
	}

	fmt.Printf("%.2f USD is %.2fEUR\n", amount, rates[EUR]*amount)

	fmt.Printf("%.2f USD is %.2fSGD\n", amount, rates[SGD]*amount)

	fmt.Printf("%.2f USD is %.2fINR\n", amount, rates[INR]*amount)

}
