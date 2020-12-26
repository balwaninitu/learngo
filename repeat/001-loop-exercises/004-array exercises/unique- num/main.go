package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	const max = 5

	var unique [max]int

	rand.Seed(time.Now().UnixNano())

loop:
	for i := 0; i < max; {

		n := rand.Intn(max) + 1

		for _, u := range unique {

			if u == n {

				continue loop
			}
		}

		unique[i] = n
		i++

	}

	fmt.Println(unique)

}
