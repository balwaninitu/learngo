package main

import "fmt"

func main() {

	i := 5

	incr(&i)

	fmt.Println(&i)
}

func incr(x *int) {

	*x++

}
