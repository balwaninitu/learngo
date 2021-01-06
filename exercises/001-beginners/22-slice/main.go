package main

import "fmt"

func main() {

	s := []string{"a", "b", "c", "d", "e"}

	s = s[:]

	x := s[len(s)-4]

	fmt.Println(s)

	fmt.Println(x)

	l := len(s)

	fmt.Println(l)

	m := s[l/2]

	fmt.Println(m)

}
