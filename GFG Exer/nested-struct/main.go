package main

import "fmt"

type emp struct {
	id int

	dept string

	details
}

type details struct {
	name string

	salary int

	post string
}

func (d details) checkmethod(ppost string) string {

	return d.post + ppost
}

func main() {

	emp1 := emp{

		id: 1234,

		dept: "IT",

		details: details{

			name: "Ana",

			salary: 5000,

			post: "developer",
		},
	}

	fmt.Println("Name -", emp1.name)

	fmt.Println("dept -", emp1.dept)

	fmt.Println("salary -", emp1.salary)

	fmt.Println("post -", emp1.post)

	fmt.Println("Post", emp1.checkmethod(" Sr"))
}
