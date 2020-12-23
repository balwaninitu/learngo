package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	if len(os.Args) != 2 {

		fmt.Println("pick a number")

		return
	}

	turns := 5

	guess, err := strconv.Atoi(os.Args[1])

	if err != nil || guess <= 0 {

		fmt.Println("Pick a positive number")

		return
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < turns; i++ {

		n := rand.Intn(guess + 1)

		if n == guess {

			fmt.Println("you win!")

			break

		}

		fmt.Println("You Lost!")

		break

	}

}
