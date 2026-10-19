package main

import (
	"fmt"
)

func calculateGrade(marks [3]int) (int, float64, string) {
	var totalMarks int
	var hasFailed bool

	for _, mark := range marks {
		totalMarks += mark

		// A mark below 40 in any subject means failure.
		if mark < 40 {
			hasFailed = true
		}
	}

	average := float64(totalMarks) / float64(len(marks))

	// A failed subject overrides the grade based on the average.
	if hasFailed {
		return totalMarks, average, "Fail"
	}

	var grade string

	if average >= 90 {
		grade = "A"
	} else if average >= 80 {
		grade = "B"
	} else if average >= 70 {
		grade = "C"
	} else if average >= 60 {
		grade = "D"
	} else {
		grade = "E"
	}

	return totalMarks, average, grade
}

func main() {
	var marks [3]int

	// Read and validate marks for each subject.
	for i := 0; i < len(marks); i++ {
		fmt.Printf("Enter marks for subject %d: ", i+1)

		_, err := fmt.Scan(&marks[i])
		if err != nil {
			fmt.Printf("Error: Please enter a valid integer for marks %d.\n", i+1)
			return
		} else if marks[i] < 0 || marks[i] > 100 {
			fmt.Println("Error: Marks must be between 0 and 100.")
			return
		}
	}

	// Calling the garadeCalculate
	totalMarks, average, grade := calculateGrade(marks)
	fmt.Println("Total Marks:", totalMarks)
	fmt.Printf("Average Marks: %.2f\n", average)
	fmt.Println("Grade:", grade)
}
