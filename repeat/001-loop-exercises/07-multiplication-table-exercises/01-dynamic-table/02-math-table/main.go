package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	const (
		validOp = "%*/+-"
	)

	args := os.Args[1:]

	l := len(args)

	if l < 1 {

		fmt.Println("Usage: [op=*/+-] [size]")

		return
	}

	if l == 1 {

		fmt.Printf("size is missin\n Usage: [op=*/+-] [size] ")
	}

	op := args[0]

	if strings.IndexAny(op, validOp) == -1 {

		fmt.Printf("Invalid operator\n valid operators one of : %s", validOp)

		return
	}

	n, err := strconv.Atoi(args[1])

	if err != nil || n < 0 {

		fmt.Printf("Error")

		return
	}

	fmt.Printf("%5s", op)

	for i := 0; i <= n; i++ {

		fmt.Printf("%5d", i)

	}

	fmt.Println()

	for i := 0; i <= n; i++ {

		fmt.Printf("%5d", i)

		for j := 0; j <= n; j++ {

			var res int

			switch op {

			case "*":

				res = i * j

			case "/":

				if j != 0 {

					res = i / j
				}

			case "%":

				if j != 0 {

					res = i % j
				}

			case "+":

				res = i + j
			case "-":

				res = i - j

			}

			fmt.Printf("%5d", res)

		}

		fmt.Println()
	}

}
