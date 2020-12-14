package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	const f = 0.3048
	const y = f / 0.9144

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)
	meters := feet * f

	yards := math.Round(feet * y)

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
	fmt.Printf("%g yards is %g meters.\n", feet, yards)

}
