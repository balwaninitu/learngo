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

		Wanna Play?

		(Provide -v flaf to see picked number)`
	)
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) < 1 {

		fmt.Printf(usage, maxTurns)

		return
	}

	var verbose bool

	if args[0] == "-v" {

		verbose = true
	}

	guess, err := strconv.Atoi(args[len(args)-1])

	if err != nil {

		fmt.Println("Not a number")

		return
	}

	if guess < 0 {

		fmt.Println("Please pick up positive number")

		return

	}

	rand.Seed(time.Now().UnixNano())

	for turn := 1; turn <= maxTurns; turn++ {

		n := rand.Intn(guess + 1)

		if verbose {

			fmt.Printf(" %d", n)
		}

		if n == guess {

			if turn == 1 {

				fmt.Println(" you are first turn winner")
			} else {

				fmt.Println(" you win")
			}

			return
		}
	}

	fmt.Println(" You lost")
}
