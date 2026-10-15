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

// Declaring arrays in Go

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

// Zero values in Go arrays.

// 1. What are zero values?
// In Go, when you decleare a variable without assigning a value, Go automatically gives it a default value called a zero value.

// For an integer array, every elements starts at 0.
func zeroValue() {
	var number [3]int

	fmt.Println(number)
}

// Output: [0 0 0]

// Why? Because numbers contains three integers, and each integer initally has the zero value 0.

// 2. Zero values for different types
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

// Initializing arrays in Go.

// 1. What is array initialization?
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

// 2. Understand the syntax
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

// Inferred Array Length in Go.

// 1. What is inferred array length?
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

// 2. Compare both declarations
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

// Accessing Array Elements Using Indexes.
// 1. What is an Index?
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

// Updating Array Elements in Go.

// 1. What does updating an array means?
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

// 2. Understand how it works
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

// Finding Array Length Using len() in Go.

// 11. What is len()?
// The build-in len()function returns the number of elements in an array.

// For example:
func lenExample() {
	numbers := [...]int{99, 54, 45, 65, 75}

	fmt.Println(len(numbers)) // Output: 5
}

// Why 5? Because array contains five elements.
// Remember: len() returns the number of elements, not the last index.

// More examples
func moreLenExample() {
	fruits := [...]string{"Apple", "Banana", "Mango", "Orange"}
	names := [...]string{"Bishan", "Mamita"}

	fmt.Println(len(fruits)) // Output: 4
	fmt.Println(len(names))  // Output: 2
}

// Even if you update an element, the array's length stays the same.
// numbers := [3]int{1, 2, 3}
// numbers[1] = 99

// fmt.Println(len(numbers))
// Output: 3

// Changing a value doesn't change the number of elements.

// 12. why is len() useful?
// You will use len() frequently when looping through arrays.

// For example:
func forLenExample() {
	numbers := [4]int{10, 20, 30, 40}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

// Don't worry if the loop is not completely clear yet. We'll study looping throught arrays in upcomming.

// Looping Through Arrays Using for

// 1. Why do we use a loop with an array?
// Suppose you have this array:
func loopWithArray() {
	numbers := [5]int{10, 20, 30, 40, 50}

	// You could print each elements individually:
	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(numbers[3])
	fmt.Println(numbers[4])

	// But imagine an array containing 100 elements. Writing 100 print statements would be inconvenient!
	// Instead, use a for loop to access each element automatically.
}

// 2. Using for with array indexes
func forWithArrayIndexes() {
	numbers := [5]int{45, 63, 74, 2, 25}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Output:
	// 45
	// 63
	// 74
	// 2
	// 25
}

// Let's understand the loop:
//
//	for i := 0; i < len(numbers); i++ {
//		fmt.Println(numbers[i])
//	}

// Code						Meaning
// i := 0				Start at index 0.
// i < len(numbers)		Continue while the index is less the 5.
// i++					Increase the index by 1
// numbers[i]			Access the elements at the current index

// 3. Follow the loop step by step
// Explore how the loop access each element.

// 10				20				30				40				50
// Index 0			Index 1			Index 2			Index 3			Index 4

// Current index: 0
// Current value: 10
// The loop prints numbers[{step}], then increases i by 1.

// Next

// Current index: 1
// Current value: 20
// The loop prints numbers[{step}], then increases i by 1.

// Next

// Current index: 2
// Current index: 30
// The loop prints numbers[{step}], then increases i by 1.

// Next

// Current index: 3
// Current value: 40
// The loop prints numbers[{step}], then increases i by 1.

// Next

// Current index: 4
// Current value: 50
// The loop prints numbers[{step}], then increases i by 1.

// Next

// Loop finished! The next index is 5, and 5 < 5 is false.

// The loop stops when i reaches 5, because the condition i < len(numbers) becomes false.
// In Go array index start at 0 not 1 so notic that the last valid index is 4, not 5.

// Looping Through Arrays Using range in Go.

// 1. What is range?
// In Go, range let's you loop through an array and get each element's index and value automatically.
// Previously, you used a regular for loop.
func previouslyLoop() {
	numbers := [3]int{1, 2, 3}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

// You can also achieve the same result more simply with range.
func rangeLoop() {
	numbers := [3]int{45, 65, 85}

	for value := range numbers {
		fmt.Println(value)
	}

	// Output:
	// 0 45
	// 1 65
	// 2 85
}

// 2. Understand the syntax
// for index, value := range numbers {
// 	fmt.Println(index, value)
// }

// Code						Meaning
// range numbers 			Iterates ove the array
// index					Receives the current element's index
// value					Receives the current element's value
// := 						Decleare's the loop variables

// For the array [10, 20, 30, 40], Go automatically provides:
//	. Index 0, value 10
//	. Index 1, value 20
//	. Index 2, value 30
//	. Index 3, value 40

// Important: The value variable contains a copy of the element's value, not a reference that automatically updates the array.

// Complete example
func rangeLoopExample() {
	names := [3]string{"Bishan", "Mamita", "Tamang"}

	for index, vlaue := range names {
		fmt.Printf("Index: %d, Value: %s\n", index, vlaue)
	}

	// Output:
	// 	Index: 0 Value: Bishan
	// Index: 1 Value: Mamita
	// Index: 2 Value: Tamang

	// Here, %d formats an integer, %s formats an string, and \n starts a new line.
}

// Using _ with range when you only need the values.

// 1. Why do we used _?
// In Go, when you use range with an array, it gives you two values:
//	. Index -> the postion of the element.
//	. Value -> the element itself.

// For example:
func blankIdentifier() {
	// numbers := [4]int{10, 20, 30, 40}

	// for index, value := range numbers {
	// 	fmt.Println(value)
	// }

	// Here, we decleare both index and value, but we only use value.
	// Go doesn't allow you to declear a local variable and leave it unused. So this code produces a compile-time error because index is decleared but not used.
}

// 2. Solution: Use _
// Go provides a special blank identifier: _.
// It let's you ignore a value you don't need.
func blankIdentifier_() {
	numbers := [4]int{10, 20, 30, 40}

	for _, value := range numbers {
		fmt.Println(value)
	}

	// Output:
	// 10
	// 20
	// 30
	// 40
}

// In this example:
//	. _ ignores the index.
//	. value received each array element.

// Remember: Use _ when you don't need a value returned by an operation.

// 3. Compare the two approaches

// When you need both index and value
//	for index, value := range numbers {
//		fmt.Println(value)
//	}

// When you olny need the value
//	for _, value := range numbers {
//			fmt.Println(value)
//	}

// Array Bounds and Out-of-Range Errors in Go.

// 1. What are array bounds?
// Array bounds refer to the valid indexes you can access an array's elements.

// For example:
// numbers := [3]int{10, 20, 30}

// This array has three elements.
// Index		Value		Valid?
// 0			10			Yes
// 1			20			Yes
// 2			30			Yes
// 3 			-			No

// Rule: If an array has length n, its vlaid indexes are from 0 to n-1.

// 2. What is an out-of-range error?
// An out-of-range error occurs when you try to access an index that doesn't exist.

// For example:
// func outOfRangeError() {
// 	numbers := [3]int{10, 20, 30}

// 	fmt.Println(numbers[3])
// }

// This program compiles, but when executed, it panics with an error similar to:
// invalid argument: index 3 out of bounds [0:3]compilerInvalidIndex

// Why?
//	. The array's length is 3.
//	. It's valid indexes are 0, 1, 2.
//	. Index 3 doesn't exist.

// Go stop normal execution of the program becuase of the invalid access.

// 3. How to avoid this error
// Use a loop condition that keeps the index below the array's length.
func avoidErro() {
	numbers := [3]int{10, 20, 30}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Output:
	// 10
	// 20
	// 30
}

// Notic that we use <, not <=.
//	. i < len(numbers) stop before index 3.
//	. i <= len(numbers) would eventually attempt to access invalid index .
