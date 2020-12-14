package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args

	if len(args) != 2 {

		fmt.Println("requires age")

		return

	}

	age, err := strconv.Atoi((args)[1])

	if err != nil || age < 0 {

		fmt.Println("Wrong age :", args[1])

	} else if age < 13 {

		fmt.Println("PG-Rated")

	} else if age > 17 {

		fmt.Println("R-Rated")

	} else {

		fmt.Println("PG-13")
	}

}
