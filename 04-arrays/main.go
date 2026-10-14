package main

import "fmt"

// 1. What is Array?
// An array is a collection of a fixed number of values of the same type.

// Imagine you want to store three exam marks.
// Without an array, you might write.
func withoutArray() {
	mark1 := 70
	mark2 := 85
	mark3 := 90

	fmt.Println(mark1)
	fmt.Println(mark2)
	fmt.Println(mark3)
}

// With an array, you can store them together:
func withArray() {
	var marks [3]int = [3]int{70, 85, 90}

	fmt.Println(marks)
}

// Here is what each part means:
// Code				Meaning
// var				Declears a variable
// marks 			Array ariable name
// [3]				Exactly three elements
// int 				Every elements stores an integer
// {70, 85, 90}		Initial values

// You can also use the shorter declaration:
func shortDeclearArray() {
	marks := [3]int{70, 85, 90}

	fmt.Println(marks)
}

// Both create an array containing three integers.

// How an array stores values
// Array: marks := [3]int{70, 85, 90}
// 			70			85			90
// 		  index 0	  index 1	  index 2
// There are 3 elements, and their indexes start at 0.

// Remember these two rules:
//	. An aaray has a fixed length.
//	. All elements in an array must have the same data types

// For example, [3] int can hold three integers, but it can't hold a string as one of its elements.
// Like this:
// marks := [3]int{"Pass", 85, 90}

// Declaring arrays

// You can declare an array using var followed by it's name, length and elements type.
var number [3]int

// Let's break it down:
//	. var - declares a variable
//	. number - the array's name
//	. [3]	- the array can hold exactly 3 elements
//	. int 	- each element stores an integer

// Example:
func declare() {
	var number [3]int

	fmt.Println(number)
}

// Output: [0 0 0]

// Why does it print [0 0 0] instead of an empty array? Go automatically initializes an integer array's elements to 0. This is called the zero value, which we will learn upcomming.

// 2. What are zero values?
// In Go, when you decleare a variable without assigning a value, Go automatically gives it a default value called a zero value.

// For an integer array, every elements starts at 0.
func zeroValue() {
	var number [3]int

	fmt.Println(number)
}

// Output: [0 0 0]

// Why? Because numbers contains three integers, and each integer initally has the zero value 0.

// 3. Zero values for different types
// Type			Zero value				Example array
// int			0						[3]int -> [0 0 0]
// float64		0						[2]float64 -> [0 0]
// string		""(empty string)		[2]string -> ["" ""]
// bool			false					[2]bool -> [false false]

// Example:
func zeroValue1() {
	var names [2]string
	var ages [2]int
	var height [2]float64
	var isStudent [2]bool

	fmt.Println(names)
	fmt.Println(ages)
	fmt.Println(height)
	fmt.Println(isStudent)
}

// Outputs:
// ["" ""]
// [0 0]
// [0 0]
// [false false]

// The exact printed representation of the string array is [ ] with empty elements; Go's formatting makes those empty strings appear between the brackets.

// 4. What is array initialization?
// Initializing an array means creating an array and assigning its initial values.

// Previously, I have learned declaring arrays:
// var number [3]int

// Output: [0 0 0]

// Now, let's initialize it with our own values:
func initialize() {
	number := [3]int{10, 20, 30}

	fmt.Println(number)
}

// Output: [10, 20, 30]

// 5. Understand the syntax
// number := [3]int{10, 20, 30}

// Part				Meaning
// number			Array name
// :=				Declares and initializes the variable
// [3]				Three elements
// int				Integer type
// [10, 20, 30]		Initial values

// You can also write the declaration this way:
// var numbers [int] = [3]int{10, 20, 30}

// Both are valid. Inside a function, the shorter := form is usually convenient.

// Example:
func initializationExample() {
	numbers := [3]int{7, 10}

	fmt.Println(numbers)
	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
}

// Output:
// [7, 10, 11]
// 7
// 10
// 11

// Remember, indexs start at 0.
//	. numbers[0] -> 7
//	. numbers[1] -> 10
//	. numbers[2] -> 11

// 6. What is inferred array length?
// Previously, i have learned to initialize array like this:
// numbers := [3]int{7, 10, 11}

// Here, you manually specify the array length as 3.

// Go, also let's you write:
func inferredArray() {
	numbers := [...]int{4, 10, 11}

	fmt.Println(numbers)
}

// The ... tells Go to count the initial values and determine the array's length automatically.
// In this example, Go counts three values, so the array's type is [3] int.

// 7. Compare both declarations
// Explicit length								Inferred length
// [3]int{1, 2, 3}								[...]int{1, 2, 3}
// You specify 3								Go counts the values
// Array length is 3							Array length is 3

// Both create an array with the same values and type.
func inferredExample() {
	studnetNames := [...]string{"Bishan", "Mamita"}

	fmt.Println(studnetNames)
	fmt.Println(len(studnetNames))
}

// Output:
// [Bishan Mamita]
// 2

// Notice that Go automatically determines the length of studentNames array is 2.

// Important: The ... syntax is used here when declearing and initializing an array. It doesn't mean the array can grow later.

// 8. What is an Index?
// An index is the postion of an elements inside an array.

// In Go, array indexes start at 0, not 1.
// For example:
func indexStart() {
	numbers := [...]int{1, 2, 3, 4, 5}

	// Array indexes and their values
	//		1				2				3				4				5
	//	index 0			index 1			index 2			index 3			index 4

	// You access an elements using the array name followed by it's index in square brackets.

	fmt.Println("First:", numbers[0])
	fmt.Println("Second:", numbers[1])
	fmt.Println("Third:", numbers[2])
	fmt.Println("Fourth:", numbers[3])
	fmt.Println("Fifth:", numbers[4])

	// Output:
	// First: 1
	// Second: 2
	// Third: 3
	// Fourth: 4
	// Fifth: 5
}

// Remember this rule:
// For an  array of length n, valid indexes range from 0 to n-1.
// For example, an array of length 5 has indexes 0, 1, 2, 4 and 5.

// 9. What does updating an array means?
// Updating an array means changing the value stored at a particular index.

// For example, you have this array:
func updatingArray() {
	numbers := [...]int{10, 20, 30}

	fmt.Println("Before:", numbers)
	// Output: [10, 20, 30]
	// Suppose you want to change 20 to 99. You can assign a new value to index 1.
	numbers[1] = 99

	fmt.Println("After:", numbers)
	// Now the array contains:
	// Output: [10, 99, 30]

	// Only the elements at index 1 changes. The other elements stay the same.
}

// 10. Understand how it works
// Before updating
//
//		10				20				30
//	Index 0			Index 1			Index 2

//					numbers[1] = 99

// After updating
//
//		10				99				30
//	Index 0			Index 1			Index 2

// Remember the difference
// . numbers[1] accesses the value at index 1.
// . numbers[1] = 99 changes the value at index 1.

// Example
func exampleUpdatingArray() {
	names := [...]string{"Mamita", "Bishan", "Jimba", "Lopchan"}

	fmt.Println("Before:", names)

	names[2] = "Lopchan"

	fmt.Println("After:", names)
}

// Output:
// Before: [Mamita Bishan Jimba Lopchan]
// After: [Mamita Bishan Lopchan Lopchan]

// The array becomes [Mamita Bishan Lopchan Lopchan].
