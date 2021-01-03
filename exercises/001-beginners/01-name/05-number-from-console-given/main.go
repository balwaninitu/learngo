package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a number: ")
	str1, _ := reader.ReadString('\n')

	// remove newline
	str1 = strings.Replace(str1, "\n", "", -1)

	// convert string variable to int variable
	num, e := strconv.Atoi(str1)
	if e != nil {
		fmt.Println("conversion error:", str1)

		return
	}

	if num >= 1 && num <= 10 {
		fmt.Println("correct")
	} else {
		fmt.Println("num not in range")
	}
}
