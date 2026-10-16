package main

import "fmt"

// Problem 6: Print all elements
// Use an indexed for loop to print each element on a separate line.
func printAllElements() {
	numbers := [5]int{10, 20, 30, 40, 50}

	// Using for loop to print each element on a separate line
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

// Problem 7: Use range
// Use range to print each element's index and value.
func useRange() {
	numbers := [5]int{10, 20, 30, 40, 50}

	// using range to print each elements' index and value
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

// Problem 8: Print only values
// Use range and _ to print only the values, not the indexes.
func onlyValue() {
	numbers := [5]int{10, 20, 30, 40, 50}

	// Using range and _ to print only the values
	for _, value := range numbers {
		fmt.Println("Value:", value)
	}
}

// Problem 9: Calculate the sum
// Write a program that calculates the sum of all elements.
func calculateSum() {
	numbers := [5]int{10, 20, 30, 40, 50}

	sum := 0

	for _, value := range numbers {
		sum = sum + value
	}

	fmt.Println("Total sum of all elements is:", sum)
}

// Problem 10: Find the largest number
// find and print the largest number
func largestNumber() {
	numbers := [6]int{10, 20, 30, 40, 50, 60}

	// Initialize largest
	largest := numbers[0]

	// Check that if value is greater then largest
	for _, value := range numbers {
		if value > largest {
			largest = value
		}
	}

	// Printing the largest number from the array
	fmt.Println("Largest Number:", largest)
}

// Calling all function in main function
func main() {
	printAllElements()
	useRange()
	onlyValue()
	calculateSum()
	largestNumber()
}
