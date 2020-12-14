package main

import "fmt"

func main() {

	counter, factor := 45, 0.5

	counter++
	counter++
	counter++
	counter++
	counter++

	factor--
	factor--

	fmt.Println(counter)
	fmt.Println(factor)
	fmt.Println("prod :", float64(counter)*factor)
}
