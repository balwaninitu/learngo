package main

import "fmt"

func main() {

	var (
		i   int
		max int
	)

	for {

		if i > max {

			break

		} else if i != 0 {

			i++

			continue

		}

		fmt.Printf("%d", i)

	}

}
