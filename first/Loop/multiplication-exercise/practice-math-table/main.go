package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	usaMsg = "Usage : [op=" + validOp + " ] [size]"

	validOp = "%*/+-"

	invalidMsg = `Invalid operator.
	Valid ops one of : ` + validOp

	invalidOp = -1

	sizemissmsg = "Size is missing"
)

func main() {

	args := os.Args[1:]

	switch l := len(args); {

	case l == 1:

		fmt.Println(sizemissmsg)

		fallthrough

	case l < 1:

		fmt.Println(usaMsg)

		return
	}

	op := args[0]

	if strings.IndexAny(op, validOp) == invalidOp {

		fmt.Println(invalidMsg)

		return

	}

	size, err := strconv.Atoi(args[1])

	if err != nil || size <= 0 {

		fmt.Println("Wrong size")

		return
	}

	fmt.Printf("%s", op)

	for i := 0; i <= size; i++ {

		fmt.Printf("%5d", i)
	}

	fmt.Println()

	for i := 0; i <= size; i++ {

		fmt.Printf("%5d", i)

		for j := 0; j <= size; j++ {

			var pr int

			switch op {

			case "*":

				pr = i * j

			case "/":

				if j != 0 {

					pr = i / j
				}

			case "%":

				if j != 0 {

					pr = i % j
				}

			case "+":

				pr = i + j

			case "-":

				pr = i - j

				fmt.Printf("%5d", pr)

				fmt.Println()

			}

		}

	}

}
