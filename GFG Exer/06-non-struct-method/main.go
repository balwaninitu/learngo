package main

import "fmt"

type val1 string

type val2 int

func (a val1) display() val1 {

	return a + "school"
}

func (b val2) display() val2 {

	return b + 45
}

func main() {

	res1 := val1("go")

	res2 := val2(45)

	fmt.Println("val1", res1.display())

	fmt.Println("val2", res2.display())

}
