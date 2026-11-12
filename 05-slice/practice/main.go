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

// Copying slices — 10 exercises

// Q14. Copy [10,20,30] into another slice using make and copy.
func Q14() {
	numbers := []int{10, 20, 30}

	// Creating destination slice using make() build-in function
	copiedSlice := make([]int, len(numbers))

	// Copy the slice from source to desination
	copy(copiedSlice, numbers)

	fmt.Printf("Orginal slice: %v\n", numbers)
	fmt.Printf("Copied slice: %v\n", copiedSlice)
}

// Q15. Copy [1,2,3,4], then change the first element of the copy to 99.
func Q15() {
	var size int
	fmt.Print("Enter the number of elements:")
	// Read the size from the user
	fmt.Scan(&size)

	// Create a slice of the size specified by the user
	numbers := make([]int, size)

	// Loop to take each number from the user
	for i := 0; i < size; i++ {
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	// Creating destination slice using make() build-in function
	copiedSlice := make([]int, len(numbers))

	// Copy the slice from source to destination
	copy(copiedSlice, numbers)

	// Changing the first elements of the copy to 99
	copiedSlice[0] = 99

	fmt.Printf("Orginal slice: %v\n", numbers)
	fmt.Printf("Copied slice: %v\n", copiedSlice)
}

// Q16. Copy [10,20,30,40,50], then change the last element of the copy to 100.
func Q16() {
	var size int
	fmt.Print("Enter the number of elements: ")

	// Read the size from the user
	fmt.Scan(&size)

	// Create a slice of the size specified by the user
	numbers := make([]int, size)

	// Loop to take each number from the user
	for i := 0; i < size; i++ {
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	// Creating destination slice using make() build-in function
	copiedSlice := make([]int, len(numbers))

	// Copy the slice from source to destination
	copy(copiedSlice, numbers)

	// Change the last element of the copy to 100
	copiedSlice[4] = 100

	fmt.Printf("Original slice:%v\n", numbers)
	fmt.Printf("Copied slice:%v\n", copiedSlice)
}

// Q17. Copy a slice and print the original length and copied length.
func Q17() {
	numbers := []int{10, 20, 30, 40}

	copied := make([]int, len(numbers))

	copy(copied, numbers)

	fmt.Println("Original slice length:", len(numbers))
	fmt.Println("Copied slice length:", len(copied))
}

// Q18. Copy [10,20,30,40,50] into a destination slice of length 3.
func Q18() {
	original := []int{10, 20, 30, 40, 50}

	copied := make([]int, 3)

	copy(copied, original)

	fmt.Println("Copied slice:", copied)

}

// Q19. Copy [10,20,30] into a destination slice of length 5.
func Q19() {
	var size int
	fmt.Print("Enter a number of elements: ")
	// Read the size from the user
	fmt.Scan(&size)

	numbers := make([]int, size)

	// Loop to take each number from the user
	for i := 0; i < size; i++ {
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	// Creating destination slice using make() build-in function
	copiedSlice := make([]int, 5)

	// Copy the slice from source to destination
	copy(copiedSlice, numbers)

	fmt.Println("Copied Slice:", copiedSlice)
	fmt.Println("Orginal slice:", numbers)
}

// Q20. Create a copy of an empty slice. Print its length and contents.
func Q20() {
	var size int

	fmt.Print("Enter the number of elements: ")

	// Read the slice size and check for invalid input
	_, err := fmt.Scan(&size)

	if err != nil {
		fmt.Println("Invalid input! Please enter an integer.")
		return
	}

	// Ensure the slice size is not negative
	if size < 0 {
		fmt.Println("Size cannot be negative.")
		return
	}

	// Create a slice with the user-specified length
	numbers := make([]int, size)

	// Take input for each element of the original slice
	for i := 0; i < size; i++ {
		fmt.Printf("Enter element %d: ", i+1)

		// Read each element and check for invalid input
		_, err := fmt.Scan(&numbers[i])

		if err != nil {
			fmt.Println("Invalid input! Please enter an integer.")
			return
		}
	}

	// Create an empty destination slice with length 0
	copiedSlice := make([]int, 0)

	// Copy elements from the original slice into the destination slice
	// No elements are copied because the destination slice has length 0
	copy(copiedSlice, numbers)

	// Print the copied slice length and contents
	fmt.Println("Copied slice length:", len(copiedSlice))
	fmt.Println("Copied slice contents:", copiedSlice)

	// Print the original slice length and contents
	fmt.Println("Original slice length:", len(numbers))
	fmt.Println("Original slice contents:", numbers)
}

// Q21. Copy a slice, append new elements to the copy, and show that the original remains unchanged.
func Q21() {
	var size int

	fmt.Print("Enter the number of elements: ")

	// Read and validate the original slice size
	_, err := fmt.Scan(&size)

	if err != nil {
		fmt.Println("Invalid input! Please enter an integer.")
		return
	}

	if size <= 0 {
		fmt.Println("Size must be greater than zero.")
		return
	}

	// Create the original slice
	numbers := make([]int, size)

	// Take input for the original slice
	for i := 0; i < size; i++ {
		fmt.Printf("Enter element %d: ", i+1)

		_, err := fmt.Scan(&numbers[i])

		if err != nil {
			fmt.Println("Invalid input! Please enter an integer.")
			return
		}
	}

	// Ask how many new elements the user wants to append
	var appendSize int

	fmt.Print("\nHow many elements do you want to append? ")

	_, err = fmt.Scan(&appendSize)

	if err != nil {
		fmt.Println("Invalid input! Please enter an integer.")
		return
	}

	if appendSize < 0 {
		fmt.Println("Number of additional elements cannot be negative.")
		return
	}

	// Create a slice to store additional elements
	newElements := make([]int, appendSize)

	// Take input for additional elements
	for i := 0; i < appendSize; i++ {
		fmt.Printf("Enter new element %d: ", i+1)

		_, err := fmt.Scan(&newElements[i])

		if err != nil {
			fmt.Println("Invalid input! Please enter an integer.")
			return
		}
	}

	// Create a destination slice with the same length as the original
	copied := make([]int, len(numbers))

	// Copy elements from the original slice
	copy(copied, numbers)

	// Append all new elements to the copied slice
	copied = append(copied, newElements...)

	// Display the original and copied slices
	fmt.Println("\nOriginal slice:", numbers)
	fmt.Println("Copied slice:", copied)
}

// Q22. Copy [5,10,15,20,25]. Modify the copied slice by doubling every element using a loop.
func Q22() {
	var size int

	fmt.Print("Enter the number of elements: ")

	// Read the slice size and check for invalid input
	_, err := fmt.Scan(&size)

	if err != nil {
		fmt.Println("Invalid input! Please enter an integer.")
		return
	}

	// Prevent the user from entering a negative slice size
	if size < 0 {
		fmt.Println("Size cannot be negative.")
		return
	}

	// Create a slice with the user-specified size
	numbers := make([]int, size)

	// Read each element from the user
	for i := 0; i < size; i++ {
		fmt.Printf("Enter number %d: ", i+1)

		// Store the user's input at the current index
		_, err := fmt.Scan(&numbers[i])

		if err != nil {
			fmt.Println("Invalid input! Please enter an integer.")
			return
		}
	}

	// Create a new slice with the same length as the original slice
	copied := make([]int, len(numbers))

	// Copy all elements from the original slice into the new slice
	copy(copied, numbers)

	// Double each element in the copied slice using its index
	for i := range copied {
		copied[i] = copied[i] * 2
	}

	// Display the original and modified slices
	fmt.Println("Original slice:", numbers)
	fmt.Println("Copied slice:", copied)
}

// Q23. Create a slice of student marks, copy it, increase every copied mark by 5, and print both slices.
func Q23() {
	var size int

	fmt.Print("Enter the number of students: ")

	// Read the number of students and check for invalid input
	_, err := fmt.Scan(&size)

	if err != nil {
		fmt.Println("Invalid input! Please enter an integer.")
		return
	}

	// Prevent negative slice sizes
	if size < 0 {
		fmt.Println("Size cannot be negative.")
		return
	}

	// Create a slice to store student marks
	studentMarks := make([]int, size)

	// Read each student's marks
	for i := 0; i < size; i++ {
		fmt.Printf("Enter marks for student %d: ", i+1)

		_, err := fmt.Scan(&studentMarks[i])

		if err != nil {
			fmt.Println("Invalid input! Please enter an integer.")
			return
		}
	}

	// Create a new slice with the same length as the original
	copiedStudentMarks := make([]int, len(studentMarks))

	// Copy all student marks into the new slice
	copy(copiedStudentMarks, studentMarks)

	// Increase each copied mark by 5
	for i := range copiedStudentMarks {
		copiedStudentMarks[i] = copiedStudentMarks[i] + 5
	}

	// Display the original and modified marks
	fmt.Printf("Original student marks: %v\n", studentMarks)
	fmt.Printf("Copied student marks: %v\n", copiedStudentMarks)
}

// Q24. Create a slice containing five numbers. Find the largest and smallest number using a loop.
func Q24() {
	numbers := []int{45, 12, 89, 34, 7, 65}

	if len(numbers) == 0 {
		fmt.Println("Slice is empty")
		return
	}

	largest := numbers[0]
	smallest := numbers[0]

	for _, n := range numbers {
		if n > largest {
			largest = n
		}
		if n < smallest {
			smallest = n
		}
	}

	fmt.Println("Largest:", largest)
	fmt.Println("Smallest:", smallest)
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
	Q14()
	Q15()
	Q16()
	Q17()
	Q18()
	Q19()
	Q20()
	Q21()
	Q22()
	Q23()
	Q24()
}
