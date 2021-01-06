package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("When is thursday?")

	today := time.Now().Weekday()

	switch time.Thursday {

	case today + 0:

		fmt.Println("Today")

	case today + 1:

		fmt.Println("Tomorrow")

	case today + 2:

		fmt.Println("In two days")

	default:

		fmt.Println("Too far away")

	}

}
