package main

import "fmt"

func main() {

	//games := []string{}

	//fmt.Println("Length", len(games))

	//fmt.Println("capacity", cap(games))

	//games = append(games, "pacman", "mario", "tetris", "doom")

	//fmt.Printf("length: %d\n Capacity: %d", len(games), cap(games))

	games := []string{"pacman", "mario", "tetris", "doom"}

	for i := 0; i <= len(games); i++ {

		s := games[:i]

		fmt.Printf("games [%d] length %d\n capacity: %d", i, len(s), cap(s))
	}

	fmt.Println()

	zero := games[:0]

	fmt.Printf("len: %d\n cap: %d\n", len(games), cap(games))

	fmt.Printf("len: %d\n cap: %d\n", len(zero), cap(zero))

	fmt.Println()

	for _, v := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {

		zero = append(zero, v)

		fmt.Printf("len: %d\n cap: %d\n", len(zero), cap(zero))

	}

	for n := range zero {

		s := zero[:n]

		fmt.Printf("zero[%d:]\n len: %d\n cap:%d\n", n, len(s), cap(s))

	}

	zero[0] = "command & conquer"
	games[0] = "red alert"

	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)

	_ = games[:cap(games)+1]
	// or:
	_ = games[:5]

}
