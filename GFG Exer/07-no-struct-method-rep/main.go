package main

import "fmt"

type data int

func (a data) add(b data) data {

	return a + b
}

func main() {

	res1 := data(6)

	res2 := data(8)

	final := res1.add(res2)

	fmt.Println("data", final)

}
