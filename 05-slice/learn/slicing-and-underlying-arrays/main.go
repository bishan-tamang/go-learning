package main

import "fmt"

// 11. Slicing Expressions - numbers[start:end]
// 11.1 What is a slicing expression?
// A slicing expression let's you take a portion of an existing slice without creating a completely separate copy of it's elements.

// For example:
func slicingExpression() {
	numbers := []int{10, 20, 30, 40, 50}

	part := numbers[1:4]

	fmt.Println(part)
}

// Output: [20, 30 40]
// Why? Because numbers[1:4] selects elements starting at index 1 and stoping before index 4.

// 11.2 Understanding the indexes
// Orginal slice: numbers
//		Index 0		Index 1		Index 2		Index 3		Index 4
//		10			20			30			40			50

// numbers[1:4] selects [20, 30, 40]
// Index 1 is included. Index 4 is excluded.

// The syntax is:
// slice[start:end]
//	. start: the index where selection begins, included.
//	. end: the index where selection stops, excluded.

// Easy rule Include the start Exclude the end.

// 11.3 More example:
// Assume we have this slice:

// Expression		Result					Explanation
// numbers[0:2]		[10, 20]				Index 0 up to, but not including, 2
// numbers[1:3]		[20, 30]				Index 1 and 2
// numbers[2:5]		[30, 40,50]				Index 2 through 4
// numbers[0:5]		[10, 20, 30, 40, 50]	All elements

// 11.4 You can leave out the start or end.
// Go allows shorter slicing expressions.

// Omit the start:
// numbers[:3]

// This is equivalent to numbers[0:3].
// Output: [10, 20, 30]

// Omit the end:
// numbers[2:]

// This selects from index 2 through the last element.
// Ouput: [30,40, 50]

// Omit both
// numbers[:]

// This select all elements in the slice.
// Output: [10, 20, 30, 40, 50]

// 11.5 Important: Slicing does not copy the elements
// This is one of the most important concept in Go Slices.
func noCopyElementInSlice() {
	numbers := []int{10, 20, 30, 40, 50}

	parts := numbers[1:3]
	parts[0] = 99

	fmt.Println(numbers)
	fmt.Println(parts)
}

// Output:
// [10, 99, 30, 40, 50]
// [99, 30]

// Why did numbers change?
// Because parts and numbers share the same underlaying array. Changing an elements throught part also changes that elements in numbers.

// 12. Slice Bounds in GO.
// In last lession, you learned how to select part of a slice using numbers[start:end].
// Now let's learn which indexes are valid and how to avoid index out of range errors.

func main() {
	slicingExpression()
	noCopyElementInSlice()
}
