package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]

	l := len(args)

	if l == 0 {

		fmt.Println(" Please give me up to 5 numbers.")

		return
	}

	if l > 5 {

		fmt.Println("Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...")

		return
	}

	var num [5]float64

	for i, v := range args {

		n, err := strconv.ParseFloat(v, 64)

		if err != nil {

			continue
		}

		num[i] = n

	}

	for range num {

		for i, v := range num {

			if i < len(num)-1 && v > num[i+1] {

				num[i], num[i+1] = num[i+1], num[i]
			}

		}
	}

	fmt.Println(num)
}
