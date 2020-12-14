package main

import "fmt"

func main() {

	work := 10

	if work > 8 {

		fmt.Println("good")

	} else if work == 8 {

		fmt.Println("average")

	} else if work == 7 {

		fmt.Println("meh..")

	} else {

		fmt.Println("bad")
	}

}
