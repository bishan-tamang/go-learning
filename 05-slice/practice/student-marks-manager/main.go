package main

import "fmt"

// Variable: pass mark used across the program
const passMark = 40

// ---------- Functions ----------

// Add a mark (returns the updated slice)
func addMark(marks []int, mark int) []int {
	if mark < 0 || mark > 100 { // operators: || and > <
		fmt.Println("Marks must be between 0 and 100.")
		return marks
	}
	return append(marks, mark)
}

// Calculate average
func average(marks []int) float64 {
	if len(marks) == 0 {
		return 0
	}
	sum := 0
	for _, m := range marks {
		sum += m
	}
	return float64(sum) / float64(len(marks))
}

// Find highest mark
func highest(marks []int) int {
	if len(marks) == 0 {
		return 0
	}
	max := marks[0]
	for _, m := range marks {
		if m > max {
			max = m
		}
	}
	return max
}

// Filter passing marks into a new slice
func passing(marks []int) []int {
	var result []int
	for _, m := range marks {
		if m >= passMark {
			result = append(result, m)
		}
	}
	return result
}

// Show all marks with a grade (switch)
func displayAllMarks(marks []int) {
	if len(marks) == 0 {
		fmt.Println("No marks stored yet.")
		return
	}
	for i, m := range marks {
		var grade string
		switch {
		case m >= 80:
			grade = "A"
		case m >= 60:
			grade = "B"
		case m >= passMark:
			grade = "C"
		default:
			grade = "F"
		}
		fmt.Printf("%d. %d  (Grade %s)\n", i+1, m, grade)
	}
}

// ---------- Main ----------

func main() {
	marks := []int{85, 72, 38, 91, 55, 40, 29}
	running := true

	for running {
		fmt.Println("\n=== Student Marks Manager ===")
		fmt.Println("1. Show all marks")
		fmt.Println("2. Add mark")
		fmt.Println("3. Average & highest")
		fmt.Println("4. Show passing marks")
		fmt.Println("5. Exit")
		fmt.Print("Choose: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			displayAllMarks(marks)
		case 2:
			var mark int
			fmt.Print("Enter mark: ")
			fmt.Scan(&mark)
			marks = addMark(marks, mark)
		case 3:
			if len(marks) == 0 {
				fmt.Println("No marks stored yet.")
			} else {
				fmt.Printf("Average: %.2f\n", average(marks))
				fmt.Println("Highest:", highest(marks))
			}
		case 4:
			p := passing(marks)
			fmt.Printf("Passing students: %d of %d\n", len(p), len(marks))
			fmt.Println(p)
		case 5:
			running = false
			fmt.Println("Goodbye!")
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
