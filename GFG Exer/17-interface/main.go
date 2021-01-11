package main

import "fmt"

type rect interface {
	Rarea() float64
	volume() float64
}

type val struct {
	length float64
	height float64
}

func (v val) Rarea() float64 {

	return v.length * v.height
}

func (v val) volume() float64 {

	return v.length * v.height * 3
}

func main() {

	var r rect

	r = val{23.6, 12.5}

	fmt.Println("Area", r.Rarea())

	fmt.Println("Volume", r.volume())

}
