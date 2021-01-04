package main

import "fmt"

func vals() (int, int) {

	return 5, 9
}

func main() {

	v1, v2 := vals()

	defer fmt.Println(v1)

	fmt.Println(v2)

	v3, _ := vals()

	defer fmt.Println("v3", v3)

}
