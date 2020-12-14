package main

import "fmt"

func main() {

	//integer types
	var i int
	var i8 int
	var i16 int
	var i32 int
	var i64 int

	//float types
	var f32 float32
	var f64 float64

	//complex types
	var c64 complex64
	var c128 complex128

	// bool type
	var b bool

	//string types
	var s string
	var r rune  //also a numeric type
	var by byte //also a numeric type

	fmt.Println(
		i, i8, i16, i32, i64,
		f32, f64,
		c64, c128,
		b, r, by,
	)

	fmt.Printf("%q\n", s)
}
