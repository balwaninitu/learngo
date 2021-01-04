package main

import "fmt"

func main() {

	s := "hello world"

	s1 := s[2:6]

	l := len(s)

	s2 := s[len(s)-4:]

	fmt.Println(l)

	fmt.Println(s1)

	fmt.Println(s2)

}
