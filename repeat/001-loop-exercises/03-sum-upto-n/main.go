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

	min, err := strconv.Atoi(args[1])

	max, err2 := strconv.Atoi(args[2])

	if err != nil || err2 != nil || min >= max {

		fmt.Println("Error")

		return
	}

	var sum int

	for i := min; i <= max; i++ {

		sum += i

		fmt.Print(i)

		if i < max {

			fmt.Print(" + ")
		}

	}

	fmt.Printf(" = %d", sum)

}
