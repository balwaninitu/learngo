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

	guess1, err1 := strconv.Atoi(os.Args[1])

	guess2, err2 := strconv.Atoi(os.Args[2])

	min := guess1

	if guess1 > guess2 {

		min = guess2
	}

	if err1 != nil && err2 != nil {

		fmt.Println("pick a positive number")

		return
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < turn; i++ {

		n := rand.Intn(min + 1)

		if n == guess1 || n == guess2 {

			fmt.Println("You won!")

			break
		}

		fmt.Println("you lost")

		break

	}

}
