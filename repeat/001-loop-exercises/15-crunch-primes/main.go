package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]

	for _, a := range args {

		ar, err := strconv.Atoi(a)

		if err != nil {
			continue

		}

		if ar == 2 || ar == 3 {

			fmt.Printf("%-2d", ar)

		} else if ar%2 == 0 && ar%3 == 0 {

			break

		} else {

			fmt.Printf("%-2d", ar)
		}

	}

}
