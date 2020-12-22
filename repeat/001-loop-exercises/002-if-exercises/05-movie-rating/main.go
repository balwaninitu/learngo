package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args

	if len(args) != 2 {

		fmt.Println("Requires age")

		return
	}

	age, err := strconv.Atoi(args[1])

	if err != nil || age < 0 {

		fmt.Printf("Wrong age %d", age)

		return

	} else if age > 17 {

		fmt.Println("R-Rated")

	} else if age < 13 {

		fmt.Println("PG-Rated")

	} else {

		fmt.Println("PG-13")
	}

}
