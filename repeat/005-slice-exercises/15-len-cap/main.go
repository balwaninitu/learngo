package main

import "fmt"

func main() {

	// var games [] string

	// games := []string{}

	// games = append(games, "pacman", "mario", "tetris", "doom")

	games := []string{"pacman", "mario", "tetris", "doom"}

	fmt.Printf("len: %d cap: %d\n", len(games), cap(games))

	fmt.Println()

	for i := 0; i <= len(games); i++ {

		s := games[:i]

		fmt.Printf("games[%d]'s len : %d cap: %d\n", i, len(s), cap(s))

		fmt.Println()

	}

	zero := games[:0]

	fmt.Printf("games: len: %d cap: %d\n", len(games), cap(games))

	fmt.Printf("zero's len: %d, cap: %d\n", len(zero), cap(zero))

	fmt.Println()

	zero = append(zero, "poceman")

	// fmt.Printf("zero's len: %d, cap: %d\n", len(zero), cap(zero))

	// fmt.Println()

	for _, v := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {

		zero = append(zero, v)

		fmt.Printf("zero's len: %d, cap: %d\n", len(zero), cap(zero))

		fmt.Println()

		for n := range zero {

			s := zero[:n]

			fmt.Printf("zero[%d]'s  len: %d cap: %d \n", n, len(s), cap(s))

			fmt.Println()

			zero = zero[:cap(zero)]
			for n := range zero {
				s := zero[:n]

				fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", n, len(s), cap(s), s)

			}

		}

	}

}
