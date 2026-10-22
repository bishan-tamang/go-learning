package main

import (
	"fmt"
)

// 1. What is Slice in Go?
// A slice is flexible way to work with a sequence of elements in Go. It let's you access elements stored in an underlying array and work with a variable
// number of elements.

// Imagine you want to store the marks of three student.

// Instead of creating three separate variables:
// mark1 := 80
// mark2 := 90
// mark3 := 75

// You can use a slice:
// marks := []int{80, 90, 75}

// Now all three marks are stored in one variable called marks.

// Understand this code
//							Slice variable
//								marks
//						Provides access to the elements
//			Index 0			Index 1			Index 2
//				80				90				75

// Coneptually, the slice provides access to these elements through an underlying array.

// Here is what each part means:
//	. marks - the variable name.
//	. [] - indicates a slice type.
//	. int - the type of elements stored in the slice.
//	. {80, 90, 75}	- the initial values.

// For example:
func slice() {
	marks := []int{80, 90, 75}

	fmt.Println(marks)
}

// Output: [80, 90, 75]

// Notice that Go print all elements inside the slice.

// 2. Slices vs Array in Go
// An array has a fixed size. A slice can grow or shrink in length as you work with it.

// 2.1 What is an array?
// An array stores a fixed number of elements. It's size is part of its type.

func array() {
	numbers := [3]int{10, 20, 30}

	fmt.Println(numbers)
}

// Output: [10, 20, 30]

// Here, [3]int means an array containing exactly 3 integers.
// You can not add a fourth element to this array becasues its size is fixed.

// 2.2 What is a slice?
// A slice lets you work with a sequence of elements whose length can change.

func sliceExample() {
	numbers := []int{10, 20, 30}

	numbers = append(numbers, 40)

	fmt.Println(numbers)
}

// Output: [10, 20, 30, 40]

// Notice that we added 40 using append(). The slice now has four elements.

// 2.3 Compare them side by side
// Array: [3]int
//		10		20		30
// Fixed size: 3 elements. You cannot add another elements to the array.

// Slice: []int
//		10		20		30		40
// Flexible length: you can add elements and save the returned slice.

// Features							Array					Slice
// Example type						[3]int					[]int
// Size								Fixed					Length can change
// Can use append() directly?		No						Yes
// Has len()?						Yes						Yes
// Has cap()?						No						Yes

// One important detail: arrays are value, and slice are also values. Assigning an array copies its elements,
// while assigning a slice copies its descriptor and shares access to the underlying array.

// 3 Declaring slices in Go
// 3.1 What does declaring a slice mean?
// Declaring a slice means creating a variable that can hold a slice value.

// The simplest syntax is:
// var numbers []int

// Let's understand each part:
// code				Meaning
// var				Declares a variable
// number			The variable's name
// []				Indicates a slice
// int				The type of elements the slice can hold

// This declares numbers as a slice of integers.

// 3.2 What is inside the slice initially?
// Consider this program:
func sliceInitially() {
	var numbers []int

	fmt.Println(numbers)
	fmt.Println(numbers == nil)
	fmt.Println(len(numbers))
}

// Output:
// []
// true
// 0

// Why?
//
//	. numbers prints as [] because it contains no elements.
//	. numbers == nil print true becasue its initial value is nil.
//	. len(numbers) print 0 because it has no elements.

// Important: A declare slice is nil by default. You can still use certain operations on it, including len() and append().

// 3.3 Can you add elements to a declared slice?
// Yes! This is one of the useful features of slices.
func addElementSlice() {
	var numbers []int

	fmt.Println("Before adding elements on numbers slice:", numbers)
	// Output: []

	numbers = append(numbers, 10)
	numbers = append(numbers, 20)
	numbers = append(numbers, 30)

	fmt.Println("After adding elements on numbers slice:", numbers)
	// Output: [10, 20, 30]
}

// Step by Stpe:
// var numbers []int			->	[]
// append(numbers, 10)			->	[10]
// append(numbers, 20)			->	[10, 20]
// append(numbers, 30)			->	[10, 20, 30]

// Notice this line:
// numbers = append(numbers, 10)
// We assign the result of append() back to numbers. This is important because append() returns the resulting slice.

// 3.4 Another way to declare a slice
// You can also declare a slice using the short declearation syntax with an inital value:
// numbers := []int{10, 20, 30}

// Compare the two approaches:
// Declare a nil slice
// var numbers []int

// Declare and initialize a slice
// numbers := []int{10, 20, 30}

// The second example of creates a slice with three initial elements. The first declare a slice with no elements.
// One small Go rule := is used inside function to decleare new variables. At package level, use var instead.

// 4. Nil Slice vs Empty Slice in Go
// This is an important Go Concept. Both nil slice and empty slice contain zero elements, but they are not exactly the same.

// 4.1 What is a nil slice?
// A nil slice is a slice whose value is nil.
// Example: var numbers []int

// Here, a numbers is declared but has not been initilized with a slice value.
// Let's check it:
func nilSlice() {
	var numbers []int

	fmt.Println(numbers)
	fmt.Println(numbers == nil)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	// Output:
	// []
	// true
	// 0
	// 0
}

// Understand the results:
//	. numbers prints as []
//	. numbers == nil is true.
//	. len(number) is 0.
//	. cap(number) is 0.

// 4.2 What is an empty slice?
// An empty slice is a slice with zero elements that is not nil.
// You can create one using an empty slice literal:
// numbers := []int{}

// Example:
func emptySlice() {
	numbers := []int{}

	fmt.Println(numbers)
	fmt.Println(numbers == nil)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	// Output:
	// []
	// false
	// 0
	// 0
}

// The important difference is that numbers == nil return false.

// 4.3 Compare nil and empty slices
// Nil slice
// Declaration: var numbers []int
//			0			0			true
//			Length		Capacity	Is nil?

// Empty slice
// Declaration: numbers := []int{}
//			0			0			false
//			Length		Capacity	Is nil?

// Both slices can grow using append():
func bothSlice() {
	var a []int
	b := []int{}

	a = append(a, 10)
	b = append(b, 10)

	fmt.Println(a) // Output: [10]
	fmt.Println(b) // Output: [10]
}

// You don't need to initialize a nil slice before appending elements.

// 4.4 When should you use each?
//	. Nil slice: var numbers []int -> a good default when you don't have any elements yet.
//	. Empty slice: numbers := []int{} -> useful when you specifically want a non-nil slice with zero elements.
//	. Both: len(numbers) == 0 works for either one.

// In backend developement, this distinction can matter when returning JSON. For example, a nil slice of integers its typically encoded as null,
// while a non-nil empty slice is enceded as [] by Go's standard encoding/json package.

// 5. Initializing slice in Go
// In this lession, you will learn how to create a slice with initial values using a slice literal.

// 5.1 What does initializing means?
// Initializing menas giving a variable its starting value.

// For example:
// numbers := []int{10, 20, 30}
// This create s slice named numbers containing three integers.

// Let's break it down:
// Code					Meaning
// numbers				Variable name
// := 					Declares and initializes a variable
// []int				Slice of integer
// {10, 20, 30}			initial elements

// The complete type of numbers is []int.

// 5.2 Complete Go Program
func initilizedSlice() {
	numbers := []int{10, 20, 30}

	fmt.Println(numbers)
	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])

	// Output:
	// [10, 20, 30]
	// 10
	// 20
	// 30
}

// Remember that numbers[0] accesses the first element, not the second.

// 5.3 You can initialize other types too
// A slice can store different types of values, depending on its declared element type.

func differentDataType() {
	ages := []int{18, 20, 25}
	names := []string{"Bishan", "Mamita"}
	prices := []float64{10.5, 20.75, 30.0}
	passed := []bool{true, false, true}

	fmt.Println(ages)
	fmt.Println(names)
	fmt.Println(prices)
	fmt.Println(passed)

	// Output:
	// [18, 20, 25]
	// [Bishan Mamita]
	// [10.5, 20.75, 30.0]
	// [ture, false, true]
}

// . []int stores integers.
// . []string stores strings.
// . []float64 sotres floating-point numbers.
// . []bool stores boolean values.

// Each slice must contain elements compatible with its element type.

// 5.4 Updating an initial value
// You can change an element after initializing the slice.
func updateSlice() {
	numbers := []int{10, 20, 30}
	numbers[1] = 99

	fmt.Println(numbers) // Output: [10, 99, 30]
}

// Here, numbers[1] = 99 replaces the value at index 1. It doesn't add a new element.

// 6. Creating Slices with make() in Go
// In this lession, you will learn how to create slice using Go's build-in make() function.

// 6.1 What is make()?
// The make() function creates and initializes a slice with a specificed length and,  optionally, a capacity.

// Basic Syntax
// numbers := make([]int, 3)

// This create a slice of integer with length 3. Each element initially conatins 0, the zero value for int.
// Output:
// [0, 0, 0]

// 6.2 Understand the syntax
// numbers := make([]int, 3)

// Part				Meaning
// numbers 			Variable name
// make() 			Create and initializes the slice
// []int			Slice of integers
// 3 				Initial length of slice

// The slice has three elements, and you can access, them using indexes 0, 1, 2.

// Complete Example
func makeFunction() {
	numbers := make([]int, 3)

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	// Output:
	// [0, 0, 0]
	// 3
	// 3
}

// Why is the capacity also 3? When you call make([]int, 3) without specifying a capacity, Go sets the capacity equal to the length.

// 6.4 Specifying both length and capacity
// You can also provide a third argument:
// numbers := make([]int, 3, 5)

// This means:
//	. Length: 3 - three elements currently belong to the slice.
//	. Capacity: 5 - the slice can grow to five elements before append() needs to allocate a larger backing array, assuming no other changes.

// Visual Example
//		0		0		0		-		-
//		0		1		2		3		4
// First 3 positions: slice length
// last 2 postions: extra capacity

// The last two positions are not part of the slice's currently length, so you cannot access them directly with numbers[3] or numbers[4].
// You can aapend() elements  to make them part of the slice.

// 6.5 Changeing elements and adding elements
func changeElementsAddingElements() {
	numbers := make([]int, 3, 5)

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers = append(numbers, 40)

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	// Output:
	// [10, 20, 30, 40]
	// 4
	// 5
}

// Notice the difference:
//	. numbers[0] = 10 updates an exixting element.
//	. append(numbers, 40) adds new element and returns the resulting slice.

// 6.6 make() vs slice literals
// Slice literal									make()
// numbers := []int{10, 20, 30}						numbers := make([]int, 3)
// Start with the values you provide				Start with zero values
// Length is inferred from the elements				You specify the length
// Cannot specify capacity in the literal itself	Can specify both lenght and capacity

// Both are useful. Choose that one that fits the task.

// 7. Accessing and Updating Elements in Go Slice.
// In this lession, you'll learn how to read, change and work with individual elements inside a slice using indexes.

// 7.1 Accessing Elements
// Suppose you have this slice:
// numbers := []int{10, 20, 30, 40, 50}

// Each elements has an index. In Go, indexing start at 0, not 1.
// Slice: numbers
// Index 0		Index 1		Index 2		Index 3		Index 4
// 10			20			30			40			50

// You can access an element by writing the slice name followed by its index in square brackets.
// fmt.Println(numbers[0])	// 10
// fmt.Println(numbers[1])	// 20
// fmt.Println(numbers[2])	// 30
// fmt.Println(numbers[3])	// 40
// fmt.Println(numbers[4])	// 50

// For example, numbers[2] give you 30, because index 2 is the third element.

// 7.2 Updating elements
// To change an element, assign a new value to its index.

func updatingElements() {
	numbers := []int{10, 20, 30}
	fmt.Println("Before updating elements:", numbers)
	// Output: [10, 20, 30]

	numbers[1] = 99
	fmt.Println("After updating elements:", numbers)
	// Output: [10, 99, 30]
}

// What happend?
//	. The orginal slice was [10, 20, 30].
//	. numbers[1] = 99 replace the value at index 1.
//	. The updated slice become [10, 99, 30]

// Important: Updating an elements doen't increase the slice's length. It replaces an exixting value.

// 7.3 Getting user input into a slice.
// You can update slice elements using input from a user.
func userInput() {
	numbers := make([]int, 3)

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Enter a number %d:", i+1)
		fmt.Scan(&numbers[i])
	}

	fmt.Println("Numbers:", numbers)
	fmt.Println("Show the last index:", len(numbers)-1)
}

// Input: 10, 20, 30
// Output:
// Numbers: [10, 20, 30]

// Here, make([]int, 3) create three elements and &numbers[i] allows fmt.Scan to wirte the entered value into each element.
// Notice that i + 1 is used in the prompt so that users can see numbers 1, 2 and 3, even through Go indexes the slice from 0.

// 7.4 What happend if you see an invalid index?
// Consider:
// numbers := []int{10, 20, 30}
// fmt.Println(numbers[3])

// This cause a runtime panic, because the valid indexes are 0, 1 and 2.
// Remember:
//
//	. First index: 0
//	. Last valid index: len(numbers) - 1
//	. For an empty slice, there is no valid element index.

// 8. Length and Capacity in Go slices
// In this lession, you'll learn two important build-in functions: len() and cap().
// These concepts are essential for understanding how slice work, especially when you start using append().

// 8.1 What is len()?
// len() returns the number of elements currently in a slice.

// For example:
// numbers := []int{10, 20, 30}

// fmt.Println(len(numbers))
// Output: 3

// Why? Because the slice currently conatins three elements: 10, 20, 30

// 8.2 What is cap()?
// cap() returns the capacity of a slice - the number of elements that can fit from the start of that slice to the end of its underlying array.

// Let's create a slice with make():
func capSlice() {
	numbers := make([]int, 3, 5)

	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	// Output:
	// 3
	// 5
}

// Here:
//	. len(numbes) is 3 because the slice currently conatins three elements.
//	. cap(numbers) is 5 because the slice has room for up to five elements before it needs a larger backing array, assuming no other changes.

// 8.3 Understand the difference visually
// make([]int, 3, 5)
//		0		0		0		-		-
//		0		1		2		3		4
// Length = 3(currently elements)
// Capacity = 5(availabe backing-array space)

// The final two positions are not yet accessible using numbers[3] or numbers[4], because the slice's length is only three.

// 8.4 What happends when you see append()?
// Let's see how the length changes when you add an element.
func addElement() {
	numbers := make([]int, 3, 5)

	fmt.Println("Before:", len(numbers), cap(numbers))

	numbers = append(numbers, 40)

	fmt.Println("After", len(numbers), cap(numbers))
	fmt.Println(numbers)
}

// Output:
// Before: 3 5
// After: 4 5
// [0 0 0 40]

// Notice:
//	. Length increased from 3 to 4.
//	. Capacity remained 5.
// 	. The three orginal elements were initialized to zero.

// If you append beyond the availabe capacity, Go allocates a larger backing array when needed.
// The new capacity is determined by GO's growth strategy, so don't assume it will always double.

// 9. Adding elements with append() in GO
// In this lession, you'll learn how to add elements to a slice, add multiple elements at once, and understand why you should save the result returned by append().

// 9.1 What is append()?
// append() build-in Go function that adds elements to a slice and returns the resulting slice.

// Basic Syntax
// numbers = append(numbers, 40)
// Here. 40 is added to the slice, and the returned slice is assigned back to numbers.

// 9.2 Adding one element
// Let's start with a simple example.
func appendOneElement() {
	numbers := []int{10, 20, 30}
	fmt.Println("Before append:", numbers)
	// Output: [10, 20, 30]

	numbers = append(numbers, 40)
	fmt.Println("After append:", numbers)
	// Output: [10, 20, 30, 40]
}

// Step by Step:
// Step 1: Orginal slice
//			10			20			30
//			    append(numbers, 40)
// Step 2: Resulting slice
// 10	20	30	40

// 9.3 Why do we write numbers = append(numbers, 40)?
// This is very important Go Concept.
// The append() function returns a slice. When you write:
// numbers = append(numbers, 40)
// you save the returned slice back into numbers.

// Why is that necessary? Because if the existing backing array doesn't have enough capacity, Go allocates a new, larger backing array.
// The returned slice then refers to the appropriate array.

// Remember: Always use the returned slice, usually by assigning it back to the slice variable.

// 9.4 Adding multiple elements
// You can add multiple elements in one append() call
func multipleAppend() {
	numbers := []int{10, 20}
	fmt.Println("Before multiple append:", numbers)
	//Output:
	// Before multiple append: [10, 20]

	numbers = append(numbers, 30, 40, 50)
	fmt.Println("After multiple append:", numbers)
	// Output:
	// After multiple append: [10, 20, 30, 40, 50]
}

// You can also combine two slice using ...
func comibneSlice() {
	first := []int{10, 20}
	second := []int{30, 40}

	fmt.Println("Before combine in first:", first)   // Output: Before combine in first: [10, 20]
	fmt.Println("Before combine in second:", second) // Output: Before combine in second: [30, 40]

	first = append(first, second...)

	fmt.Println("After combine:", first) // Output: After combine in first: [10, 20, 30, 40]
}

// The ... expands the elements of second so that append() adds them individually.

// 9.5 Appending to a nil slice
// Remeber Topic 3 where we decleared a slice without initial values?
// You can append directly to it:
func appendNill() {
	var numbers []int
	fmt.Println("Before append:", numbers) // Output: []

	numbers = append(numbers, 10)
	numbers = append(numbers, 20)

	fmt.Println("After append:", numbers) // Output: [10, 20]
}

// You don't need to make() first.

// 10. Looping Through Slices with for and range in Go.
// You'll learn how to access every element in a slice without writing numbers[0] numbers[1], numbers[2] and so on.

// 10.1 What is looping?
// A loop repeats a block of code. In Go, you can use for with range to go through every elements of a slice.

// For example:
func loopingThroughSlices() {
	numbers := []int{10, 20, 30, 40}

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// Output:
	// 0 10
	// 1 20
	// 2 30
	// 3 40
}

// 10.2 How does range work?
// The statement for index, value := range numbers give you two things for each element:
// index - the position of the element, starting at 0.
// value - the value stored at that postion.

// Understanding the slice
//		Index 0		Index 1		Index 1		Index 3
//		10			20			30			40

// range vistis the elements in order, from index 0 through index 3.

// 10.3 Using only the value.
// If you only need the element values, you can ignore the index using _.
func onlyValue() {
	numbers := []int{10, 20, 30, 40}

	for _, value := range numbers {
		fmt.Println(value)
	}

	// Output:
	// 10
	// 20
	// 30
	// 40
}

// Remember _ from our earlier lesson? It's the blank identifier. It tell Go that you don't need that returned value.

// 10.4 Using only the index
// If you only need the index, you can omit the value:
func onlyIndex() {
	numbers := []int{10, 20, 30}

	for index := range numbers {
		fmt.Println(index)
	}

	// Output:
	// 0
	// 1
	// 2
}

// 10.5 Updating slice elements inside a loop
// Be careful! This code does update the original slice:
func updateSliceElements() {
	numbers := []int{10, 20, 30}

	for _, value := range numbers {
		value = value * 2
	}

	fmt.Println(numbers) // Output: [10, 20, 30]
}

// Why? The value variable recevies a copy of each element.

// To update the actual slice, use the index:
func actualUpdateSlice() {
	numbers := []int{10, 20, 30}

	for index := range numbers {
		numbers[index] = numbers[index] * 2
	}

	fmt.Println(numbers) // Output: [20, 40, 60]
}

// Key rule: Use numbers[index] when you want to modify the original slice elements.

func main() {
	slice()
	array()
	sliceExample()
	sliceInitially()
	addElementSlice()
	nilSlice()
	emptySlice()
	bothSlice()
	initilizedSlice()
	differentDataType()
	updateSlice()
	makeFunction()
	changeElementsAddingElements()
	updatingElements()
	userInput()
	capSlice()
	addElement()
	appendOneElement()
	multipleAppend()
	comibneSlice()
	appendNill()
	loopingThroughSlices()
	onlyValue()
	onlyIndex()
	updateSliceElements()
	actualUpdateSlice()
}
