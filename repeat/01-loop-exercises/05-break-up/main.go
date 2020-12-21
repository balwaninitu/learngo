package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args

	if len(args) != 3 {

		fmt.Println("give two number min & max")

		return
	}

	min, err1 := strconv.Atoi(args[1])

	max, err2 := strconv.Atoi(args[2])

	if err1 != nil || err2 != nil || min >= max {

		fmt.Println("Invalid number")

		return
	}

	var (
		i   = min
		sum int
	)

	for {

		if i > max {
			break

		} else if i%2 != 0 {
			i++
			continue
		}

		fmt.Print(i)

		if i < max-1 {

			fmt.Print(" + ")
		}

		sum += i
		i++

	}

	fmt.Printf(" = %d", sum)
}
