package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	turn := 5

	guess, err := strconv.Atoi(os.Args[1])

	if err != nil {

		fmt.Println("pick a positive number")

		return
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < turn; i++ {

		n := rand.Intn(guess + 1)

		if n == guess {

			fmt.Println("you win")

			break

		}

		fmt.Println("you lost")

		break

	}

}
