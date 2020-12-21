package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args

	if len(args) != 3 {

		fmt.Println("Give two number min & max")

		return
	}

	min, err1 := strconv.Atoi(args[1])

	max, err2 := strconv.Atoi(args[2])

	if err1 != nil || err2 != nil || min >= max {

		fmt.Println("Invalid number")

		return
	}

	var sum int

	for i := min; i <= max; i++ {

		if i%2 != 0 {

			continue
		}

		sum += i

		fmt.Print(i)

		if i < max-1 {

			fmt.Print(" + ")
		}
	}

	fmt.Printf(" = %d", sum)

}
