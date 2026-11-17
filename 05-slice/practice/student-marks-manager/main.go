package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const passMark = 40

// ---------- Input Helper ----------

func readInt(reader *bufio.Reader, message string) (int, error) {
	fmt.Print(message)

	input, err := reader.ReadString('\n')
	if err != nil && len(input) == 0 {
		return 0, err
	}

	input = strings.TrimSpace(input)

	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("please enter a valid integer")
	}

	return value, nil
}

// ---------- Functions ----------

// Display all marks with grades
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

		fmt.Printf("%d. %d (Grade %s)\n", i+1, m, grade)
	}
}

// Add a mark
func addMark(marks []int, mark int) []int {
	if mark < 0 || mark > 100 {
		fmt.Println("Marks must be between 0 and 100.")
		return marks
	}

	return append(marks, mark)
}

// Remove the first matching mark
func removeMark(marks []int, target int) ([]int, bool) {
	for i, m := range marks {
		if m == target {
			marks = append(marks[:i], marks[i+1:]...)
			return marks, true
		}
	}

	return marks, false
}

// Update the first matching mark
func updateMark(marks []int, oldMark, newMark int) ([]int, bool) {
	if newMark < 0 || newMark > 100 {
		fmt.Println("Marks must be between 0 and 100.")
		return marks, false
	}

	for i, m := range marks {
		if m == oldMark {
			marks[i] = newMark
			return marks, true
		}
	}

	return marks, false
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

	highestMark := marks[0]

	for _, m := range marks {
		if m > highestMark {
			highestMark = m
		}
	}

	return highestMark
}

// Find lowest mark
func lowest(marks []int) int {
	if len(marks) == 0 {
		return 0
	}

	lowestMark := marks[0]

	for _, m := range marks {
		if m < lowestMark {
			lowestMark = m
		}
	}

	return lowestMark
}

// Filter passing marks
func passing(marks []int) []int {
	var result []int

	for _, m := range marks {
		if m >= passMark {
			result = append(result, m)
		}
	}

	return result
}

// ---------- Main ----------

func main() {
	reader := bufio.NewReader(os.Stdin)

	marks := []int{85, 72, 38, 91, 55, 85, 40, 29}

	running := true

	for running {
		fmt.Println("\n=== Student Marks Manager ===")
		fmt.Println("1. Show all marks")
		fmt.Println("2. Add mark")
		fmt.Println("3. Remove mark")
		fmt.Println("4. Average & highest")
		fmt.Println("5. Show passing marks")
		fmt.Println("6. Update mark")
		fmt.Println("7. Show lowest mark")
		fmt.Println("8. Exit")

		choice, err := readInt(reader, "Choose: ")

		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		switch choice {

		case 1:
			displayAllMarks(marks)

		case 2:
			mark, err := readInt(reader, "Enter mark: ")

			if err != nil {
				fmt.Println("Invalid input:", err)
				continue
			}

			marks = addMark(marks, mark)

		case 3:
			displayAllMarks(marks)

			target, err := readInt(reader, "Enter mark to remove: ")

			if err != nil {
				fmt.Println("Invalid input:", err)
				continue
			}

			var removed bool
			marks, removed = removeMark(marks, target)

			if removed {
				fmt.Println("Mark removed successfully!")
			} else {
				fmt.Println("Mark not found!")
			}

		case 4:
			if len(marks) == 0 {
				fmt.Println("No marks stored yet.")
				continue
			}

			fmt.Printf("Average: %.2f\n", average(marks))
			fmt.Println("Highest:", highest(marks))

		case 5:
			p := passing(marks)

			fmt.Printf("Passing students: %d of %d\n", len(p), len(marks))

			if len(p) == 0 {
				fmt.Println("No students passed.")
			} else {
				fmt.Println("Passing marks:", p)
			}

		case 6:
			displayAllMarks(marks)

			oldMark, err := readInt(reader, "Enter mark to update: ")

			if err != nil {
				fmt.Println("Invalid input:", err)
				continue
			}

			newMark, err := readInt(reader, "Enter new mark: ")

			if err != nil {
				fmt.Println("Invalid input:", err)
				continue
			}

			var updated bool
			marks, updated = updateMark(marks, oldMark, newMark)

			if updated {
				fmt.Println("Mark updated successfully!")
			} else {
				fmt.Println("Mark not found or update failed!")
			}

		case 7:
			if len(marks) == 0 {
				fmt.Println("No marks stored yet.")
			} else {
				fmt.Println("Lowest mark:", lowest(marks))
			}

		case 8:
			running = false
			fmt.Println("Goodbye!")

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
