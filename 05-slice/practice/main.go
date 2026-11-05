package main

import "fmt"

// Q1. Create a slice containing 5, 10, 15, 20, 25. Print the first element, the last element, and the length.
func Q1() {
	// Creating a slice containing 5, 10, 15, 20, 25
	numbers := []int{5, 10, 15, 20, 25}

	// Printing the first element
	fmt.Printf("First element: %d\n", numbers[0])

	// Printing the second element
	fmt.Printf("Last element: %d\n", numbers[4])

	// Printing the length of slice
	fmt.Printf("Total length of slice: %d\n", len(numbers))
}

// Q2. Given []int{10, 20, 30, 40}, use a loop to calculate and print the total.
func Q2() {
	// Creating a slice
	numbers := []int{10, 20, 30, 40}

	total := 0

	// Calculate total elements of slice
	for _, num := range numbers {
		total = num + total
	}

	fmt.Printf("Total: %d\n", total)
}

// Q3. Start with []int{1, 2, 3}. Append 4 and 5, then print the slice and its length.
func Q3() {
	// Creating slice
	numbers := []int{1, 2, 3}

	// Append 4 and 5 in numbers slice
	numbers = append(numbers, 4, 5)

	// Printing the slice after append
	fmt.Printf("After append slice: %d\n", numbers)
	// Printing its length after append
	fmt.Printf("Its length after append: %d\n", len(numbers))
}

func main() {
	Q1()
	Q2()
	Q3()
}
