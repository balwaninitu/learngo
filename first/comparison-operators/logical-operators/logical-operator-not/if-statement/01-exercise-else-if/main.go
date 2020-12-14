package main

import "fmt"

func main() {

	age := 70

	if age > 60 {

		fmt.Println("Getting older")

	} else if age > 30 {

		fmt.Println("wiser")

	} else if age > 20 {

		fmt.Println("Adulthood")

	} else if age > 10 {

		fmt.Println("Young blood")

	} else {

		fmt.Println("booting up")
	}

}
