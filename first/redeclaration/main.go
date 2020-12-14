package main

import "fmt"

func main() {
	var safe bool
	safe, speed := true, 50
	fmt.Println(safe, speed)

	name := "Nikola"
	name, age := "Marie", 20

	fmt.Println(name, age)

	name, birth := "Albert", 1879
	fmt.Println(name, birth)

}
