package main

import "fmt"

func main() {

	slice1 := []string{"go", "school", "evry", "day"}

	slice2 := []string{"try", "hard"}

	copy1 := copy(slice1, slice2)

	fmt.Println("copy", copy1)

	fmt.Println(slice1)

	fmt.Println(slice2)

}
