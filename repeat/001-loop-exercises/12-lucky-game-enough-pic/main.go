package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	guess, err := strconv.Atoi(os.Args[1])

	if err != nil {

		fmt.Println("pick number")

		return
	}

	var turn int

	if guess <= 10 {

		turn = 10

	} else {

		turn = guess
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
