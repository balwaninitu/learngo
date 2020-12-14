package main

import "fmt"

func main() {

	const (
		monday    = iota + 1
		tuesday   = iota
		wednesday = iota
		thursday  = iota
		friday    = iota
		saturday  = iota
		sunday    = iota
	)
	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

}
