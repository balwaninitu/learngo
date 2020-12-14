package main

import "fmt"

func main() {

	var (
		name   string
		age    int
		famous bool
	)

	name = "Newton"
	age = 84
	famous = true

	fmt.Println(name, age, famous)

	name = "somebody"
	age = 20
	famous = false

	fmt.Println(name, age, famous)

	var prevName string
	prevName = name

	name = "Einstein"

	fmt.Println("previous name:", prevName)
	fmt.Println("current name:", name)
}
