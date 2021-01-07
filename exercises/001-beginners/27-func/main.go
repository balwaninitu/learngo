package main

import "fmt"

func isEven(num int) bool {

	if num%2 == 0 {
		return true

	}

	return false

}

func main() {

	x := 34

	y := 45

	ans1 := isEven(x)

	ans2 := isEven(y)

	fmt.Println(ans1, ans2)

}
