package main

import "fmt"

func main() {

	defer fmt.Println("Hello")
	defer fmt.Println("!")
	fmt.Println("World")

	// defer function or method call arguments evaluate instantly,
	// but they execute until the nearby functions returns

}
