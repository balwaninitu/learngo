package main

import (
	"fmt"
)

func main() {

	// names := []string{"lina", "mina", "dica", "pica", "nica"}

	// fmt.Println("1st step", names)

	names := make([]string, 5)

	fmt.Println("1st step", names)

	names = append(names, "einstein", "tesla", "aristo")

	fmt.Println("2nd step", names)

	names = append([]string{"einstein", "tesla", "aristo"}, names...)

	fmt.Println("3rd step", names)

	moreNames := [...]string{"plato", "khayyam", "ptolemy"}

	copy(names[:2], moreNames[:2])

	fmt.Println("4th step", names)

	clone := make([]string, 3, 5)

	copy(names[:2], clone[len(clone)-3:])

	clone = append(clone, names[:2]...)

	//sliced := clone[1:4:4]

}
