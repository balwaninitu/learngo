package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for {

		var s string

		switch rand.Intn(5) {

		case 0:
			s = "\\"

		case 1:
			s = "|"
		case 2:
			s = "/"

		case 3:
			s = "-"

		case 4:
			s = "+"
		}

		fmt.Printf("\r%s Please Wait! coming soon!", s)
		time.Sleep(time.Millisecond * 500)
	}

}
