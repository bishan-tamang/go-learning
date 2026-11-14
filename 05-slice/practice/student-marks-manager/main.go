package main

import "fmt"

func findAverage(studentMarks []int) float64 {
	total := 0

	for _, num := range studentMarks {
		total = total + num
	}

	average := float64(total) / float64(len(studentMarks))

	return average
}

func findHighestMarks(studentMarks []int) int {
	highestMarks := studentMarks[0]

	for _, num := range studentMarks {
		if num > highestMarks {
			highestMarks = num
		}
	}

	return highestMarks
}

func findPassStudnet(studentMarks []int) []int {
	var passStudent []int

	for _, num := range studentMarks {
		if num >= 40 {
			passStudent = append(passStudent, num)
		}
	}

	return passStudent
}

func addStudentMarks(studentMarks []int) []int {
	var addStudentMarksSize int
	fmt.Print("How many student marks do you want to add? ")

	if _, err := fmt.Scan(&addStudentMarksSize); err != nil || addStudentMarksSize <= 0 {
		fmt.Println("Invalid input! Please enter an integer.")
		return studentMarks
	}

	newStudentMarks := make([]int, addStudentMarksSize)

	for i := 0; i < addStudentMarksSize; i++ {
		fmt.Printf("Enter new studnet marks %d: ", i+1)

		if _, err := fmt.Scan(&newStudentMarks[i]); err != nil {
			fmt.Println("Invalid input! Please enter an integer.")
			return studentMarks
		}
	}

	studentMarks = append(studentMarks, newStudentMarks...)

	return studentMarks
}

func main() {
	var size int

	fmt.Print("Enter number of elements for student marks: ")
	if _, err := fmt.Scan(&size); err != nil || size <= 0 {
		fmt.Println("Invalid input! Please enter an integer.")
		return
	}

	// Creating slice for student marks
	studentMarks := make([]int, size)
	// Read each elements from user
	for i := 0; i < size; i++ {
		fmt.Printf("Enter a student %d mark: ", i+1)
		if _, err := fmt.Scan(&studentMarks[i]); err != nil {
			fmt.Println("Invalid input! Please enter an elements.")
			return
		}

		if studentMarks[i] > 100 || studentMarks[i] < 0 {
			fmt.Println("Invalid marks")
			return
		}
	}

	averageMarks := findAverage(studentMarks)
	highestMarks := findHighestMarks(studentMarks)
	passStudent := findPassStudnet(studentMarks)
	addStudentMarks := addStudentMarks(studentMarks)

	// Diplsay the result
	fmt.Printf("Average marks: %v\n", averageMarks)
	fmt.Printf("Highest marks: %d\n", highestMarks)
	fmt.Printf("Pass studnet filter: %v\n", passStudent)
	fmt.Printf("Add studnet marks: %v\n", addStudentMarks)
}
