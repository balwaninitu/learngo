package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 3 {

		fmt.Println("Give me two no.")

		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil || min >= max {

		fmt.Println("Wrong number")

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

			fmt.Print(" +  ")
		}

		fmt.Printf(" =%d\n", sum)
	}

}
