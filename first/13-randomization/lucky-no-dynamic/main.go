package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	const (
		maxTurns = 5

		usage = `Welcome to Luck Number Game!!

		The programme will pick up %d random numbers.
		Your mission is to guess one those numbers.

		The greater your numbers is, the harder it gets.

		Wanna Play?`
	)
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 1 {

		fmt.Printf(usage, maxTurns)

		return
	}

	guess, err := strconv.Atoi(args[0])

	if err != nil {

		fmt.Println("Not a number")

		return
	}

	if guess < 0 {

		fmt.Println("Please pick up positive number")

		return
	}

	min := 10

	if guess > min {

		min = guess
	}

	for turn := (maxTurns + guess/4); turn > 0; turn-- {

		n := rand.Intn(guess + 1)

		if n == guess {

			fmt.Println("You win")

			return
		}
	}

	fmt.Println("You Lost....Try again!!")
}
