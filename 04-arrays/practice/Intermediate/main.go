package main

import "fmt"

// Problem 11: Find the smallest number
// Find the smallest element in: numbers := [6]int{25, 10, 85, 40, 15, 60}
func smallestNumber() {
	numbers := [6]int{25, 10, 85, 40, 15, 60}

	// Initialize smallest number as 0 index
	smallest := numbers[0]

	// Check that if value is smallest number smallest
	for _, value := range numbers {
		if value < smallest {
			smallest = value
		}
	}

	// Printing the smallest number from the numbers arrays.
	fmt.Println("Smallest Number:", smallest)
}

// Problem 12: Search for a number
// Ask the user to enter a number. Search for it in an array and print whether it was found or not.
func searchNumber(number int) {
	numbers := [5]int{10, 20, 30, 40, 50}
	found := false

	for _, value := range numbers {
		if number == value {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}
}

// Problem 13: Calculate the average
// Calculate the average of all the values in an integer array. Convert the sum to float64 when dividing if you want a decimal result.
func calculateAverage() {
	numbers := [4]int{10, 15, 20, 25}

	// Initialize sum for sum of numbers array
	sum := 0

	// Calculate the sum of all value
	for _, value := range numbers {
		sum = sum + value
	}

	// Finding the avearge of numbers array and converting int to float64
	average := float64(sum) / 4

	// Printing the avearge of numbers arrays
	fmt.Println("Average of all values:", average)
}

// Problem 14: Copy an array
// Create an array a, copy it into b, change one element in b, and print both arrays. Demonstrate that the original remains unchanged.
func copyArray() {
	// Creating Orginal array and print orginal array
	a := [3]int{7, 10, 11}
	fmt.Println("Orginal Array:", a)

	// Copy a array into b array and change one elemet in b and printing copied array.
	b := a
	b[1] = 9
	fmt.Println("Copied Array:", b)
}

// Problem 15: Pass an array to a function
// Write a function named changeFirst that accepts a [3]int array and changes its first element. Call the function from main() and print the original array afterward.
func changeFirst(a [3]int) {
	a[0] = 40
	fmt.Println("Copied:", a)
}

func main() {
	smallestNumber()

	var number int
	fmt.Print("Enter a number:")
	fmt.Scan(&number)

	searchNumber(number)
	calculateAverage()
	copyArray()

	a := [3]int{10, 20, 30}
	fmt.Println("Orginal:", a)
	changeFirst(a)
}
