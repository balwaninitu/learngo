package main

import "fmt"

func area(l, w int) int {

	ar := l * w

	return ar
}

func main() {

	fmt.Printf("Area is %d", area(15, 5))

}
