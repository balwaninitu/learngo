package main

import "fmt"

func main() {

	fmt.Printf("Type of %d is %[1]T\n", 3)
	fmt.Printf("Type of %.2f is %[1]T\n", 3.14)
	fmt.Printf("Type of %s is %[1]T\n", "hello")
	fmt.Printf("type of %t is %[1]T\n", true)
}
