package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 3 {

		fmt.Println("Give me two number")

		return
	}
	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil || min >= max {

		fmt.Println("Wrong no.")

		return
	}

	var (
		sum int
		i   = min
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

	fmt.Printf(" =%d\n", sum)

}
