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
		maxTurns = 8

		Usage = `  Welcome to Lucky number Game!
		
		This game will pick %d random number.
		Your mission is to guess one those number.
		
		The greater your number is, the harder it gets.
		
		Wanna Play?
		
		If yes please pick a positive number`
	)

	args := os.Args[1:]

	if len(args) != 1 {

		fmt.Printf(Usage, maxTurns)

		return

	}

	guess, err := strconv.Atoi(args[0])

	if err != nil {

		fmt.Println("Wrong number")

		return
	}

	if guess < 0 {

		fmt.Println("Pick a positive number")

		return
	}

	rand.Seed(time.Now().UnixNano())

	for turn := 0; turn < maxTurns; turn++ {

		n := rand.Intn(guess + 1)

		if n == guess && turn == 1 {

			fmt.Println("Congatulations!! You are First Turn Winner!")

			return

		} else if n == guess && turn != 1 {

			fmt.Println("Congatulations!! You Won!!")

			return

		}

	}

	fmt.Println("Sorry you lost..Try again!!")
}
