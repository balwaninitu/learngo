package main

import "fmt"

func main() {

	const min = 1 + 1
	const version = "2.0.1" + "-beta"
	const pi = 3.14 * min
	const debug = !true

	fmt.Println(min, version, pi, debug)

}
