package main

import "fmt"

// Problem 1: Declare an array
// Declare an array named numbers that stores 5 integers. Print the array.
func declareArray() {
	// Decleare an array
	var numbers [5]int

	// Printing the array without assigning value
	fmt.Println(numbers)
}

// Problem 2: Initialize an array
// Create an array containing 10, 20, 30, 40, 50. Print it.
func initializeArray() {
	// Declare and initilize an array
	numbers := [5]int{10, 20, 30, 40, 50}

	fmt.Println(numbers)
}

// Problem 3: Access elements
// Given: numbers := [5]int{10, 20, 30, 40, 50}
func accessElements() {
	numbers := [5]int{10, 20, 30, 40, 50}

	// Accessing first, third and fifth elements.
	fmt.Println(numbers[0])
	fmt.Println(numbers[2])
	fmt.Println(numbers[4])
}

// Problem 4: Update elements
// Given the same array, change the second element to 99 and the last element to 100. Print the updated array.
func updateElements() {
	numbers := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Before update:", numbers)

	// Change the second element to 99 and the last elemet to 100
	numbers[1] = 99
	numbers[4] = 100

	fmt.Println("After update:", numbers)
}

// Problem 5: Find the length
// Create an array of 6 integers and use len() to print how many elements it contains.
func findLength() {
	numbers := [6]int{10, 20, 30, 40, 50, 60}

	// Print how many elements it contains inside the numbers array
	fmt.Println(len(numbers))
}

func main() {
	declareArray()
	initializeArray()
	accessElements()
	updateElements()
	findLength()
}
