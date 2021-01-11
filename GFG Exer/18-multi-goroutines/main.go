package main

import (
	"fmt"
	"time"
)

func name() {

	arr1 := [3]string{"asha", "misha", "tisha"}

	for i := 0; i < 3; i++ {

		time.Sleep(150 * time.Millisecond)

		fmt.Printf("%s\n", arr1[i])
	}

}

func id() {

	arr2 := [3]int{4, 5, 2}

	for v := 0; v < 3; v++ {

		time.Sleep(300 * time.Millisecond)

		fmt.Printf("%d\n", arr2[v])
	}
}

func main() {

	fmt.Println("Main func starts")

	go name()

	go id()

	time.Sleep(3500 * time.Millisecond)

	fmt.Println("\n..Main go-routine End..!")

}
