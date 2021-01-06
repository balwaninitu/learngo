package main

import (
	"fmt"
	"time"
)

func main() {

	m := time.Now().Month()

	switch time.September {

	case m + 0:

		fmt.Println("Birth month")

	case m + 1:

		fmt.Println("Valentine month")

	case m + 8:

		fmt.Println("special month")

	default:

		fmt.Println("Not a birth month")

	}

}
