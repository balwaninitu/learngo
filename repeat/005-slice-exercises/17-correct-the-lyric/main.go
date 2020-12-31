package main

import (
	"fmt"
	"strings"
)

func main() {

	lyric := strings.Fields(`all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay`)

	lyric = append([]string{"yesterday"}, lyric...)

	fmt.Printf("%s\n", lyric)
}
