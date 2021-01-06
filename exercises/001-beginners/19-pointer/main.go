package main

import "fmt"

func main() {

	i, j := 25, 78

	p := &i

	fmt.Println(*p)

	*p = 34

	fmt.Println(i)

	p = &j

	fmt.Println(*p)

	fmt.Println(j)

	*p = *p / 2

	fmt.Println(*p)

	fmt.Println(j)

	fmt.Println(i)

	fmt.Printf("%T", p)

}
