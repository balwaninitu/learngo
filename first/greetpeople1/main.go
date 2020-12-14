package main

import (
	"fmt"

	"os"
)

func main() {

	name := os.Args[1]

	fmt.Println("Hello there!", name, "here!")

	name, age := "Riddham", 8

	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("I love to go school!")

}
