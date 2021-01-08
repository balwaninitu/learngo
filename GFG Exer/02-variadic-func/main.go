package main

import (
	"fmt"
	"strings"
)

func joinStr(element ...string) string {

	return strings.Join(element, "-")
}

func main() {

	school := []string{"temasek", "primary", "school"}

	fmt.Println(joinStr("go", "school", "everyday"))

	fmt.Println(joinStr(school...))

}
