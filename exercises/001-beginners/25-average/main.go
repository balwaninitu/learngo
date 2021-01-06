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

	class1 := []int{86, 80, 75, 82, 80, 70}
	class2 := []int{40, 55, 60, 48, 50}
	class3 := []int{90, 85, 30, 86, 78, 75, 54}

	var sum1, sum2, sum3 float64

	for _, v := range class1 {

		sum1 += float64(v)

	}

	ave1 := sum1 / float64(len(class1))

	fmt.Printf("Average mark of class 1 is %.2f\n", ave1)

	for _, v := range class2 {

		sum2 += float64(v)

	}

	ave2 := sum2 / float64(len(class2))

	fmt.Printf("Average mark of class 2 is %.2ff\n", ave2)

	for _, v := range class3 {

		sum3 += float64(v)

	}

	ave3 := sum3 / float64(len(class3))

	fmt.Printf("Average mark of class 3 is %.2ff\n", ave3)

	avg := ave1 + ave2 + ave3

	fmt.Printf("Average mark for all 3 classes  is %.2f\n", avg/3)

}
