package main

import "fmt"

func main() {
	// DO NOT TOUCH THIS
	isSphere, radius := true, 25

	if isSphere && radius >= 200 {
		fmt.Println("It's a big sphere.")

	} else {
		fmt.Println("I don't know.")
	}
}
