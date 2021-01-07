// class_1 = [86, 80, 75, 82, 80, 70]
// class_2 = [40, 55, 60, 48, 50]
// class_3 = [90, 85, 30, 86, 78, 75, 54]

// marks_list = [class_1, class_2, class_3]

// a) Using a nested loop, calculate and display the average mark for each class. Your answer is expected to be in this format

// Average Mark for Class 1 is 78.83

// Average Mark for Class 2 is 50.60

// Average Mark for Class 3 is 71.14

// b) Using a nested loop, calculate and display the average mark for all 3 classes and the highest and lowest marks.

// Your answer is expected to be in this format

// Average Mark for all 3 Classes is 68.00

// The Highest Mark is 90

// The Lowest Mark is 30
package main

import "fmt"

func main() {

	s2 := [][]int{
		{86, 80, 75, 82, 80, 70},
		{40, 55, 60, 48, 50},
		{90, 85, 30, 86, 78, 75, 54},
	}

	rows := len(s2)

	var totalSum int

	var totalLen int

	for i := 0; i < rows; i++ {
		sum := 0
		for j := 0; j < len(s2[i]); j++ {
			sum = sum + s2[i][j]
		}

		totalSum = totalSum + sum

		totalLen = totalLen + len(s2[i])

		avg1 := float64(sum) / float64(len(s2[i]))

		fmt.Printf("Average of class %d is %.2f\n", i, avg1)

	}

	totalAvg := totalSum / totalLen

	fmt.Printf("Average of all class is %d\n", totalAvg)

}
