package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}

	fmt.Printf("%q\n\n", ships)

	from, to := 0, len(ships)

	switch num := os.Args[1:]; len(num) {

	default:

		fallthrough

	case 0:

		fmt.Println("Provide only the [starting] and [stopping] positions")

		return

	case 2:
		to, _ = strconv.Atoi(num[1])
		fallthrough

	case 1:
		from, _ = strconv.Atoi(num[0])

	}

	l := len(ships)

	if from < 0 || from > l || from > to || to > l {

		fmt.Println("wrong position")

		return

	}

	fmt.Println(ships[from:to])

}
