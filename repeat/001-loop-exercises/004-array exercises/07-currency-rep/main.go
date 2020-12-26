package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	const (
		SGD = iota
		INR
		Eur
		Jyp
	)

	args := os.Args[1:]

	if len(args) != 1 {

		fmt.Println("Please provide the amount to be converted.")

		return
	}

	rates := [...]float64{

		SGD: 3.6,
		INR: 5.8,
		Eur: 60.3,
		Jyp: 80.5,
	}

	n, err := strconv.ParseFloat(args[0], 64)

	if err != nil {

		fmt.Println("Invalid amount. It should be a number.")
	}

	fmt.Printf("%.2f USD is %.2f SGD\n", n, rates[SGD]*n)
	fmt.Printf("%.2f USD is %.2f INR\n", n, rates[INR]*n)
	fmt.Printf("%.2f USD is %.2f Eur\n", n, rates[Eur]*n)
	fmt.Printf("%.2f USD is %.2f Jyp\n", n, rates[Jyp]*n)

}
