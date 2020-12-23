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

	guess, err := strconv.Atoi(os.Args[1])

	if err != nil || guess <= 0 {

		fmt.Println("pick a positive number")

		return
	}

	rand.Seed(time.Now().UnixNano())

	turn := 5

	for i := 1; i < turn; i++ {

		n := rand.Intn(guess + 1)

		if n == guess && i == 1 {

			fmt.Println("Congrats! You are first turn winner")

			break

		} else if n == guess {

			fmt.Println("You win!")

			break
		}

		fmt.Println("You lost!")

		break
	}

}
