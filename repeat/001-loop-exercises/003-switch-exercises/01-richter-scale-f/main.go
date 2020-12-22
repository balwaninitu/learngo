package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {

		fmt.Println("Give me the magnitude of the earthquake")

		return
	}

	richter, err := strconv.ParseFloat(os.Args[1], 64)

	if err != nil {

		fmt.Println("I couldnt get that sorry")

		return
	}

	var res string

	switch r := richter; {

	case r < 2.0:

		res = "micro"

	case r >= 2 && r < 3:

		res = "very minor"

	case r >= 3 && r < 4:

		res = "minor"

	case r >= 4 && r < 5:

		res = "light"

	case r >= 5 && r < 6:

		res = "moderate"

	case r >= 6 && r < 7:

		res = "strong"

	case r >= 7 && r < 8:

		res = "major"

	case r >= 8 && r < 10:

		res = "great"

	default:

		res = "massive"
	}
	fmt.Printf("%.2f is %s\n", richter, res)

}
