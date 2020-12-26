package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	a := os.Args[1:]

	if len(a) == 0 {

		fmt.Println("Provide few number")

		return
	}

	num := []int{}

	for _, v := range a {

		n, err := strconv.Atoi(v)

		if err != nil {

			continue
		}

		num = append(num, n)

	}

	sort.Ints(num)

	fmt.Printf("%d", num)

}
