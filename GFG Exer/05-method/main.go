package main

import "fmt"

type student struct {
	name string
	age  int
}

type teacher struct {
	subject string
	marks   int
}

func (s student) show() {

	fmt.Println("name", s.name)

	fmt.Println("age", s.age)
}

func (t teacher) show() {

	fmt.Println("subject", t.subject)

	fmt.Println("marks", t.marks)

}

func main() {

	st := student{"Preeti", 23}

	tea := teacher{"go", 30}

	st.show()

	tea.show()

}
