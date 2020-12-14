package main

import "fmt"

func main() {

	blue := [3]int{6, 9, 3}

	red := [3]int{6, 9, 3}

	fmt.Printf("blue bookcase : %v\n", blue)

	fmt.Printf("red bookcase : %v\n", red)

	fmt.Println("are they equal?", blue == red)

}
