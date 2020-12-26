package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]

oLoop:

	for _, a := range args {

		n, err := strconv.Atoi(a)

		if err != nil {

			continue
		}

		switch {

		case n == 2 || n == 3:

			fmt.Print(n, " ")

			continue

		case n%2 == 0 || n%3 == 0 || n <= 1:

			continue
		}

		for i, w := 5, 2; i*i <= n; i++ {

			if n%i == 0 {

				continue oLoop
			}

			i += w

			w = 6 - w
		}

		fmt.Print(n, " ")

	}

	fmt.Println()

}
