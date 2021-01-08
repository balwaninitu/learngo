package main

import "fmt"

func name(a, b int) (int, int, int) {

	return a + b, a - b, a * b
}

func main() {

	fmt.Println(name(4, 8))

}
