package main

import (
	"fmt"
)

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

// Slicing a slice — 10 exercises

// Q4. Given [10,20,30,40,50], extract elements from index 1 to 3.
func Q4() {
	numbers := []int{10, 20, 30, 40, 50}

	// extract elements from index 1 to 3
	num := numbers[1:3]

	fmt.Printf("Extract elements are: %d\n", num)
}

// Q5. Extract the first 3 elements from [5,10,15,20,25].
func Q5() {
	numbers := []int{5, 10, 15, 20, 25}

	// Extract the first 3 elements from [5,10,15,20,25].
	num := numbers[0:2]

	fmt.Printf("The first 3 elements are: %d\n", num)
}

// Q6. Extract elements from index 2 to the end of [10,20,30,40,50].
func Q6() {
	numbers := []int{10, 20, 30, 40, 50}

	// Extract elements from index 2 to the end of [10, 20, 30, 40, 50]
	num := numbers[2:]

	fmt.Printf("Extract elements are: %d\n", num)
}

// Q7. Extract the entire slice using slicing syntax.
func Q7() {
	numbers := []int{10, 20, 30, 40}

	// Extract the entire slice using slicing syntax.
	num := numbers[:]

	fmt.Printf("Extract the entire slice: %v\n", num)
}

// Q8. Extract the last 3 elements from [1,2,3,4,5,6].
func Q8() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	// Extract the last 3 elements
	num := numbers[3:]

	fmt.Printf("Extract the last 3 elements: %v\n", num)
}

// Q9. 	Extract the middle elements from [10,20,30,40,50,60].
func Q9() {
	numbers := []int{10, 20, 30, 40, 50, 60}

	// Extract the middle elements
	num := numbers[2:4]

	fmt.Printf("Extract the middle elements are: %v\n", num)
}

// Q10. Given [5,10,15,20,25], create a slice containing all elements except the first and last.
func Q10() {
	numbers := []int{5, 10, 15, 20, 25}

	// Creating a slice containing all elements excepts the first and last.
	num := numbers[1:4]

	fmt.Printf("Extract the all elements except the first and last: %v\n", num)
}

// Q11. Create a slice from index 1 to 4, then change its first element to 99. Print both original and sliced slices.
func Q11() {
	original := []int{10, 20, 30, 40, 50}

	// Create a slice from index 1 to 4
	slicedSlice := original[1:4]

	// Changing its first element to 99
	slicedSlice[0] = 99

	fmt.Println("Original Slice:", original)
	fmt.Println("Sliced slices:", slicedSlice)
}

// Q12. Given [10,20,30,40,50,60], create two slices: first half and second half.
func Q12() {
	numbers := []int{10, 20, 30, 40, 50, 60}

	// Create two slices: first half and second half
	firstHalf := numbers[:3]
	secondHalf := numbers[3:]

	// Printing both slice
	fmt.Printf("First half: %v\n", firstHalf)
	fmt.Printf("Second half: %v\n", secondHalf)
}

// Q13. Given [1,2,3,4,5,6,7,8], extract elements at indexes 2 through 5. Calculate their sum.
func Q13() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sum := 0

	// Extract elements at indexes 2 through 5.
	num := numbers[2:6]

	// Calculate their sum
	for _, value := range num {
		sum = sum + value
	}

	fmt.Printf("Sum of extract elements is: %d\n", sum)

}

func main() {
	Q1()
	Q2()
	Q3()
	Q4()
	Q5()
	Q6()
	Q7()
	Q8()
	Q9()
	Q10()
	Q11()
	Q12()
	Q13()
}
