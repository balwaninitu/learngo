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
		JYP
	)

	currency := [...]float64{

		SGD: 35.0,
		INR: 60,
		Eur: 20,
		JYP: 30,
	}

	args := os.Args[1:]

	if len(args) != 1 {

		fmt.Println("Please provide the amount to be converted")

		return
	}

	n, err := strconv.ParseFloat(args[0], 64)

	if err != nil || n < 0 {

		fmt.Println("Invalid amount. It should be a number")

		return
	}
	fmt.Printf("%.2f usd is %.2fSGD\n", n, currency[SGD]*n)

	fmt.Printf("%.2f usd is %.2fINR\n", n, currency[INR]*n)

	fmt.Printf("%.2f usd is %.2fEur\n", n, currency[Eur]*n)

	fmt.Printf("%.2f usd is %.2fJYP\n", n, currency[JYP]*n)

}
