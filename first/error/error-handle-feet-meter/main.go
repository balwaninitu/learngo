package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	const f = 0.3048

	arg := os.Args[1]

	feet, err := strconv.ParseFloat(arg, 64)

	meters := feet * f

	if err != nil {

		fmt.Printf("%q is not number.\n", arg)

		return
	}

	fmt.Printf("%g feet is %g meters.\n", feet, meters)

}
