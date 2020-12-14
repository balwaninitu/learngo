package main

import "fmt"

func main() {

	const minsPerDay = 60 * 24

	const weekDays = 7

	fmt.Printf("there are %d minutes in %d weeks.\n", minsPerDay*weekDays*2, 2)

}
