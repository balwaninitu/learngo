package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {

		fmt.Println("Give me a year number")

		return
	}

	y, err := strconv.Atoi(os.Args[1])

	if err != nil {

		fmt.Printf("%v is not a leap year\n", os.Args[1])

		return
	}

	if y%4 == 0 && (y%400 == 0 || y%100 != 100) {

		fmt.Printf("%d is a leap year\n", y)
	} else {

		fmt.Printf("%d is not a leap year\n", y)

	}

}
