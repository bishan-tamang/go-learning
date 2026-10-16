package main

import "fmt"

// Problem 16: Reverse an array
// Reverse the order of the elements without creating another array.
func reverseArray(num [5]int) {
	left := 0
	right := len(num) - 1

	for left < right {
		// Swap the elements
		num[left], num[right] = num[right], num[left]

		left++
		right--
	}

	fmt.Println("Reversed Array:", num)
}

// Problem 17: Count even and odd numbers
// Count how many even and odd numbers are present in: numbers := [8]int{12, 7, 4, 9, 16, 3, 20, 5}
func countEvenOdd(number [8]int) (int, int) {
	// Initilized even and odd as zero
	even := 0
	odd := 0

	// Check even or odd from value
	for _, value := range number {
		if value%2 == 0 {
			even = even + 1
		} else {
			odd = odd + 1
		}
	}

	return even, odd
}

// Problem 18: Compare two arrays
// Create two arrays of type [3]int. Check whether they are equal using ==. Then change one element and check again.
func compareTwoArrays(oneArray, twoArray [3]int) string {
	if oneArray == twoArray {
		return "Both are Equal."
	}

	return "Not Equal"
}

// Problem 19: Find the second-largest distinct number
// Find the second-largest distinct value in: numbers := [6]int{10, 50, 30, 50, 20, 40}
func secondLargest(number [6]int) int {
	// Initialize with the first element of the array
	largest := number[0]
	secondLargest := number[0]
	foundSecond := false

	for _, value := range number {
		if value > largest {
			// Shift the old largest down to second largest
			secondLargest = largest
			largest = value
			foundSecond = true
		} else if value < largest {
			if !foundSecond || value > secondLargest {
				// Update second largest only
				secondLargest = value
				foundSecond = true
			}
		}
	}

	if !foundSecond {
		fmt.Println("No second-largest distinct number exists.")
		return 0
	}

	return secondLargest
}

// 20. Array Analyzer
// Build a console program that:
//
//	. Stores 5 integers in an array.
//	. Prints every element using a loop.
//	. Calculates the sum and average.
//	. Finds the largest and smallest values.
//	. Asks the user for a number and searches the array.
//	. Prints the total number of elements using len().
func arrayAnalyzer() {
	// 1. Store 5 integers in an array
	numbers := [5]int{25, 10, 85, 40, 15}

	// 2. Print every element using a loop
	fmt.Println("Array elements:")
	for _, value := range numbers {
		fmt.Println(value)
	}

	// 3. Calculate sum and average
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	average := float64(sum) / float64(len(numbers))

	fmt.Println("Sum:", sum)
	fmt.Println("Average:", average)

	// 4. Find largest and smallest values
	largest := numbers[0]
	smallest := numbers[0]

	for _, value := range numbers {
		if value > largest {
			largest = value
		}

		if value < smallest {
			smallest = value
		}
	}

	fmt.Println("Largest:", largest)
	fmt.Println("Smallest:", smallest)

	// 5. Ask the user for a number and search the array
	var target int
	fmt.Print("Enter a number to search: ")
	fmt.Scan(&target)

	found := false

	for _, value := range numbers {
		if value == target {
			found = true
			break
		}
	}

	if found {
		fmt.Println(target, "was found in the array.")
	} else {
		fmt.Println(target, "was not found in the array.")
	}

	// 6. Print total number of elements
	fmt.Println("Total elements:", len(numbers))
}

func main() {
	// 16.
	num := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Orginal:", num)
	reverseArray(num)

	// 17.
	number := [8]int{12, 7, 4, 9, 16, 3, 20, 5}
	even, odd := countEvenOdd(number)
	fmt.Printf("Total Even: %d, Total Odd: %d\n", even, odd)

	// 18.
	oneArray := [3]int{7, 10, 11}
	twoArray := [3]int{7, 10, 11}
	fmt.Println(compareTwoArrays(oneArray, twoArray))

	oneArray = [3]int{7, 8, 11}
	fmt.Println(compareTwoArrays(oneArray, twoArray))

	// 19.
	numbers := [6]int{50, 50, 30, 20, 10, 5}

	fmt.Println("Second Largest:", secondLargest(numbers))

	// 20.
	arrayAnalyzer()
}
