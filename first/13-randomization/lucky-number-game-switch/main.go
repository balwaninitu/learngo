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

		usage = `Welcome to Lucky number Game!!

		This game will pick %d random numbers.
		Your mission is to guess one of those number.

		The greater the number is, the harder it gets.

		Wanna Play?

		If yes, Please pick two positive number.

				!!Beast Of Luck!!`

		firstturnmsg = "Congratulations!! You are first turn winner"

		winmsg = "Congatulations!! You won!!"

		lostmsg = "Sorry !You have Lost...Try Luck again!"
	)

	args := os.Args[1:]

	if len(args) < 1 {

		fmt.Printf(usage, maxTurns)

		return
	}

	guess, err := strconv.Atoi(args[0])
	guess2, err2 := strconv.Atoi(args[1])

	var res string

	switch {

	case err != nil || err2 != nil:

		res = "Wrong number"

		fallthrough

	case guess < 0 || guess2 < 0:

		res = "Pick two positive number"

		fmt.Printf(res)

		return

	}

	min := guess

	if guess < guess2 {

		min = guess2
	}

	rand.Seed(time.Now().UnixNano())

	for turn := 0; turn < maxTurns; turn++ {

		n := rand.Intn(min + 1)

		if n == guess || n == guess2 {

			fmt.Println("you win")

			return
		}

	}

	fmt.Println("You lost")
}
