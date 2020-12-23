package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	guess := 6

	turn := 5

	maxTurn := (turn + guess/4)

	for i := 0; i < maxTurn; i++ {

		n := rand.Intn(guess + 1)

		if n == guess {

			fmt.Println("you win!")

			break
		}

		fmt.Println("you lost")

		break
	}

}
