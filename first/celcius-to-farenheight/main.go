package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	arg := os.Args[1]

	celcius, _ := strconv.ParseFloat(arg, 64)

	farenheit := celcius*1.8 + 32

	fmt.Printf("%g celcius is %g farenheit.\n", celcius, farenheit)

}
