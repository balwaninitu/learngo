package main

import (
	"fmt"
	"time"
)

func main() {

	age := 37

	//dob := 27

	year := time.Now().Year()

	//fmt.Println(year)

	birthyr := year - age

	fmt.Printf("Birth year :%d\n", birthyr)

}
