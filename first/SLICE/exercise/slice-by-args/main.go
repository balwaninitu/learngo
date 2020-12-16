package main

import (
	"fmt"
	"os"
)

func main() {

	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}

	if len(os.Args) == 1 {

		fmt.Printf("%q\n", ships)

		fmt.Println("Provide only the [starting] and [stopping] positions")

		return

	}

}
