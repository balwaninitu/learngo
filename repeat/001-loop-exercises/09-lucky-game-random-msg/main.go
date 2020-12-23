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

	turn := 5

	guess, err := strconv.Atoi(os.Args[1])

	if err != nil || guess < 0 {

		fmt.Println("pick a positive number")

		return
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < turn; i++ {

		n := rand.Intn(guess + 1)

		var msg string

		if n == guess {

			switch rand.Intn(3) {

			case 0:

				msg = "you won"

			case 1:

				msg = "you are awesome"

			case 2:

				msg = "you did it!"

			}

			fmt.Printf("%s", msg)

			return

		}

		switch rand.Intn(2) {

		case 0:

			msg = ("you lost")

		case 1:

			msg = ("bad luck!")
		}

		fmt.Println(msg)

		break

	}

}
