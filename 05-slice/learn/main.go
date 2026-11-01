package main

import (
	"fmt"
)

// 1. What is Slice?
// A slice in Go is a dynamically sized, flexible view into an underlying array.
// Unlike arrays which have a fixed size, slices can grow and shrink automatically, making them the standard way to handle collections of data in Go.

// For example, imagine you want to store the marks of five students.
// 1. Without slice
func WithoutSlice() {
	student1 := 70
	student2 := 85
	student3 := 90
	student4 := 65
	student5 := 80

	fmt.Println(student1) // Output: 70
	fmt.Println(student2) // Output: 85
	fmt.Println(student3) // Output: 90
	fmt.Println(student4) // Output: 65
	fmt.Println(student5) // Output: 80
}

// 2. With slice
func storeMarks() {
	marks := []int{70, 85, 90, 65, 80}

	fmt.Println(marks) // Output: [70, 85, 90, 65, 80]
}

// Let's break down the declaration
// marks := []int{70, 85, 90, 65, 80}
//	. marks - variable name.
//	. []int	- a slice that stores integer values
//	. {70, 85, 90, 65, 80} - contains the initial values used to create the slice.
//	. := declare and initializes the variable.
//	. []int{18, 20, 25} - slice literal.

// A slice literal specifies the slice type and its initial elements.

// You can store strings, floats, structs and othe types in slice too, provide the elements have the same type.
// names := []string{"Ram", "Hari", "Sita"}
// prices := []float64{100.50, 250.00, 75.25}

// Real-world usages: a list of products, users, orders, database results or task names.

// 2. Arrays vs Slices
func arraySlice() {
	// Arrays have a fixed length, while slices are fleible, and grow or shrink. This is the most important difference to understand.

	// Features				Array						Slice
	// Length				Fixed						Can grow or shrink
	// Declaration example  [3]int						[]int
	// Adding elements		Can't change length			Can use append() build-in function
	// Comman use			Fixed-size collections		Flexible collection

	numbers := [3]int{10, 20, 30} // Array
	scores := []int{10, 20, 30}   // Slice

	// Adding elements in Array
	// numbers = append(numbers, 60)

	// But change the value
	numbers[2] = 7
	fmt.Println(numbers)

	// Adding elements in Slice
	scores = append(scores, 7)
	fmt.Println(scores)

	// Notic that we wrote scores = append(scores, 7), not simply append(scores, 7).
	// Because append return a slice, and that returned slice must be used to update scores.

	// Remember: In everyday Go programming, you will often use slices whenever you need a collection whose size can change.
}

// 3. Creating slices: the four pattern you need
// Pattern A: Create a slice with values
func sliceWithValues() {
	numbers := []int{10, 20, 30}

	fmt.Println(numbers)

	// Use this when you already know the initial values.
}

// Pattern B: Declare a nil slice
func declareNilSlice() {
	var numbers []int

	fmt.Println(numbers)        // Output: []
	fmt.Println(len(numbers))   // Output: 0
	fmt.Println(numbers == nil) // Output: true
	// When you declare a slice using var without giving it a value, Go automatically assigns its zero value, which is nil. So, numbers is a nil slice.

	// This creates a nil slice. It has no elements, but you can append to it.
	numbers = append(numbers, 10)
	fmt.Println(numbers)
}

// Pattern C: Create an empty, non-nil slice
func emptySlice() {
	numbers := []int{}

	fmt.Println(numbers)        // Output: []
	fmt.Println(len(numbers))   // Output: 0
	fmt.Println(numbers == nil) // Output: false
	// Here, Go does not automatically assign the zero value nil. You explicitly create an empty slice using []int{}.
}

// Pattern D: Use make()
func makeFunction() {
	numbers := make([]int, 3)

	fmt.Println(numbers)
	fmt.Println(len(numbers))

	// The syntax is:
	// make([]int, length, capacity)

	// The capacity is optional.
	total := make([]int, 3, 5)

	fmt.Println(total)

	// Here:
	// . Length = 3
	// . Capacity = 5
	// . Initial Values = [0 0 0]

	// Important: make([]int, 3) create three existing elements. It does not merely reserve space for three elements.
}

// 4. len() and cap() - Understand this visualization
//	. len() is a built-in function that returns the number of elements in a slice, array, string, map, or channel.
//	. Capacity (cap) means how many elements a slice can hold in its underlying array starting from the slice's first element before it needs more storage.

// Example: numbers := make([]int, 3, 5)
//		10			20			30			-
//		Element		Element		Element		Unused

// Length: 3
// The last positions are availabe for growth, but they are not accessible as slice elements unitl the length increases.
func lenCap() {
	// What happend when you append?
	numbers := make([]int, 2, 3)

	// Before append
	fmt.Println("Before append:", numbers) // Ouput: [0, 0]
	fmt.Println(len(numbers))              // Output: 2
	fmt.Println(cap(numbers))              // Output: 3

	// After append
	numbers = append(numbers, 10)

	fmt.Println("After append:", numbers) // Output: [0, 0, 10]
	fmt.Println(len(numbers))             // Output: 3
	fmt.Println(cap(numbers))             // Output: 3
	// The slice now uses all three available slots.

	// What happens when you append beyond capacity?
	numbers = append(numbers, 20)

	// Now the slice needs room for four elements, but its old capacity was only three.

	// Go automatically allocates more storage when necessay and updates the slice to use it.
	// The new capaicty is at least 4, but you shouldn't assum it will be exactly 4 or double the previous capacity.

	fmt.Println(numbers)      // Output: [0 0 10 20]
	fmt.Println(len(numbers)) // Ouput: 4
	fmt.Println(cap(numbers)) // Output: 6

	// Comman Mistake:
	// numbers[4] = 30

	fmt.Println(numbers) // Output: panic: runtime error: index out of range [4] with length 4

	// Why? The slice's length is only 4, so valid indexes are 0, 1, 2, 3. Capacity does not make index 4 accessible.
	// Use append to add a new element.
}

// 5. Accessing and Updating slice elements
// Slice indexes start at zer, just like array indexes.
func accessingUpdatingSliceElements() {
	makrs := []int{70, 85, 90, 65}

	// Accessing slice elements
	fmt.Println(makrs[0]) // Output: 70
	fmt.Println(makrs[2]) // Output: 90

	// updating slice elements
	makrs[1] = 95
	makrs[0] = 75

	fmt.Println(makrs) // Output: [75 95 90 65]

	// Comman mistakes
	//	. Accessing makrs[4] when the slcie has only four elements causes a runtime panic.
	//	. Using index 1 when you mean the first element returns the second element.
	//	. Trying to assign to an index that hasn't been added to the slice causes a panic.

	// A useful safety check is:
	if len(makrs) > 0 {
		fmt.Println(makrs[0])
	}
	// This prevents a panic when accessing the first element of an empty slice.
}

// 6. Looping through slices with range
// This is one of the most frequently used patterns in Go.
// In Go, range let's you loop through a slice and access each element without manually managing the index.
// 6.1 Basic Example
func loopWithRange() {
	numbers := []int{10, 20, 30, 40}

	for index, value := range numbers {
		fmt.Println("Index", index, "Value", value)
	}

	// Output:
	// Index 0 Value 10
	// Index 1 Value 20
	// Index 2 Value 30
	// Index 3 Value 40
}

// 6.2 How does range work?
// for index, value := range numbers

// . range numbers - goes through each element of the slice.
// . index - gets the element's position, starting from 0.
// . value - gets the element's value.

// Example slice
//		Index 0		Index 1		Index 2		Index 3
//		10			20			30			40
// Each loop iteration give you the next index and value.
// Loop iteration means one time the loop runs.

// 6.3 Get only the values
// If you don't need the index, use _ or decleare only the value by using the single-variable form.
// The underscore _ means you're intentionally ignoring the index.
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

// Remember: In for index, value := range slice, Go gives you the index first and the value second.

// 7. append() - add elements to a slice.
// 7.1 append() is a build-in Go function that adds elements to a slice and return the updated slice.
// Example:
func appendExample() {
	numbers := []int{10, 20, 30}

	numbers = append(numbers, 40)

	fmt.Println(numbers) // Output: [10 20 30 40]
	// Initially, the slice contains three elements. After appending 40, it contains four.
}

// 7.2 Understand the syntax
// append(slice, element...)

// The general syntax is:
// append(slice, element1, element2, ...)
// . append() - the build-in function
// . slice - the slice you want to add elements to.
// . element1, element2 - the elements you want to add.
// . return value - the resulting slice.

// Example
func appendExample2() {
	numbers := []int{10, 20, 30}

	result := append(numbers, 40)

	fmt.Println(result) // Output: [10 20 30 40]
	// Here, result receives the slice returned by append().
}

// 7.3 Why do we write numbers = append(numbers, 40)?
// This is one of the most important things to understand.
// numbers := []int{10, 20, 30}
// numbers = append(numbers, 40)

// Let's break it down:
// Step 1 - Orginal slice
// numbers = [10, 20, 30]
// Length: 3

// Step 2 - Append element
// append(numbers, 40)
// Go returns a slice containing the orginal elements and 40.

// Step 3 - Assign the returned slice
// numbers = [10, 20, 30, 40]
// Length: 4

// Important: append() return a slice. Assigning the result back to numbers ensures you keep the updated slice, including any new underlyingstorage Go may have allocated.

// 7.4 Appending multiple elements
// You can append more than one elements in a single call.
func appendMultiple() {
	numbers := []int{10, 20, 30}
	fmt.Println("Before append:", numbers) // Ouptut: [10 20 30]

	numbers = append(numbers, 40, 50, 60)
	fmt.Println("After append:", numbers) // Output: [10 20 30 40 50 60]

	// Here three elements are added at onece.
	// Before append	Appended		After append
	// [10 20 30]		40, 50, 60		[10 20 30 40 50 60]
}

// 7.5 Appending to an empty slice
// You can append elements even when a slice initially has no elements.
func nilSlice() {
	var numbers []int

	numbers = append(numbers, 10, 20)

	fmt.Println(numbers) // Output: [10 20]
	// This works because Go allows append() to work with nil slices. Go manages the underlyingstorage as needed.
}

func appendEmptySlice() {
	numbers := []int{}

	numbers = append(numbers, 10, 20)

	fmt.Println(numbers) // Output: [10 20]
	// Both approach are valid.
}

// 7.6 What happens to length and capacity?
// Remember:
// . len() returns the numbers of elements in the slice.
// . cap() returns the capacity from the slice's starting position.

// let's see how they change:
func appendLengthCapacity() {
	number := make([]int, 2, 3)

	fmt.Println(len(number)) // OUtput: 2
	fmt.Println(cap(number)) // Output: 3

	number = append(number, 10)

	fmt.Println(len(number)) // Output: 3
	fmt.Println(cap(number)) // Output: 3

	number = append(number, 20)

	fmt.Println(len(number)) // Output: 4
	fmt.Println(cap(number)) // Output: At least 4

	// The first append uses the remaining capacity. The sceond append exceeds the old capacity, so Go automaticlly provides more storeage.
}

// 7.7 Appending one slice to another
// Suppose you have two slices:
func appendOneSliceToAnother() {
	first := []int{10, 20}
	second := []int{30, 40}

	// To append all elements from second to first, use ...
	first = append(first, second...)

	fmt.Println(first) // Output: [10 20 30 40]

	// The ... expands the elements of second so that append() receives them as individual arguments.
	// Without ..., append(first, second...) would not compile if second is a []int, becauses
	// append(frist, second) tires to append the entire slice as a single element of type []int to a slice of type []int.
}

// 7.8 Does append() always modify the orginal underlyingarray?
// No, this is an important details.

// Go may reuse the existing underlyingarray if there is enough capacity.
// If there is not enough capacity, Go alloactes a new underlyingarray.
func modifyAppend() {
	numbers := make([]int, 2, 2)
	numbers[0] = 10
	numbers[1] = 20

	numbers = append(numbers, 30)
	// The orginal capacity was 2, and the new length is 3. Go must provide more storeage.
	// For now, remember: append() handles this automatically. Always assign its returned slice back when you want to keep the result.
}

// 8. Slicing a slice: [:]
// 8.1 What does slicing s slice means?
// Slicing means selecting a portion of an existing slice.

// For example:
func slicingExample() {
	numbers := []int{10, 20, 30, 40, 50}

	part := numbers[1:4] // Output: [20, 30, 40]

	fmt.Println(part)
	// Here, numbers[1:4] selects elements starting at index 1 and ending just before index 4.
}

// 8.2 Understand the syntax
// slice[start:end]
// . start - the index where selection begins(included).
// . end - the index where selection stops(excluded).

// Example:
// Orginal Slice
//	10		[20		30		40]		50
//	0		1		2		3		4
// numbers[1:4] select indexes 1, 2 and 3.

// Result: [20, 30, 40]

// 8.3 Comman slicing patterns
// Asume this slice:
func commanSlicingPatterns() {
	numbers := []int{10, 20, 30, 40, 50}

	part1 := numbers[1:3]
	part2 := numbers[:3]
	part3 := numbers[2:]
	part4 := numbers[:]

	fmt.Println(part1) // Output: [20 30]
	fmt.Println(part2) // Output: [10 20 30]
	fmt.Println(part3) // Output: [30 40 50]
	fmt.Println(part4) // Output: [10 20 30 40 50]
}

// Expression		Result				Meaning
// numbers[1:3]		[20 30]				Index 1 upto, but not including, 3
// numbers[:3]		[10 20 30]			From the beginning to index 3
// numbers[2:]		[30 40 50]			from index 2 to the end
// numbers[:]		[10 20 30 40 50]	Select all elements

// 8.4 What does [:] mean?
// [:] means select the entire slice, from the beginning to the end.
func selectEntireSlice() {
	numbers := []int{10, 20, 30}

	all := numbers[:]

	fmt.Println(all) // Output: [10 20 30]
	// Notice that numbers[:] does not create an independent copy of the elements. The resulting slice shares the same underlyingarray.

	// For example:
	all[0] = 99

	fmt.Println(numbers) // Output: [99 20 30]
	fmt.Println(all)     // Output: [99 20 30]

	// Both slice see the changes because they refer to the same underlyingarray.
	// Remember: [:] selects all elements, but it doesn't automatically copy them into separate storage.
}

// 9. Copy() - make an independent Copy
// The build-in copy() function is one of the most important slice operations in Go.
// It's main purpose is to copy elements from one slice into another slice without making both slice depend on the same elements storage.

// 9.1 Why do we need copy()?
// First, look at this example:
func copyExample() {
	orginal := []int{10, 20, 30}

	copied := orginal

	fmt.Println("Orginal:", orginal) // Output: [10 20 30]
	fmt.Println("Copied:", copied)   // Output: [10 20 30]

	copied[0] = 99

	fmt.Println("Orginal:", orginal) // Output: [99 20 30]
	fmt.Println("Copied:", copied)   // Output: [99 20 30]

	// Why did changing copied also change orginal?
	// Because assigning one slice to another doesn't copy the underlyingelements.
	// Both slice variables refer to the same underlyingarray.

	// Before chaning copied[0]
	//			10			20			30
	//			Index 0		Index 1		Index 2
	//						| Same underlyingarray

	// Orginal							Copied
	// Both variables refer to the same elements. Updating through either slice affects the shared data.

	// This is where copy() become useful.
}

// 9.2 What is copy()?
// Go provides a buil-in function:
// copy(distination, source)
// It copies elements from the source slice into the destination slice.

// Syntax:
// copy(destinationSlice, sourceSlice)

// For Example:
func copySliceExample() {
	orginal := []int{10, 20, 30}

	copied := make([]int, len(orginal))
	// Its initial values are: [0 0 0]
	fmt.Println("Initial values:", copied) // Output: [0 0 0]
	copy(copied, orginal)

	copied[0] = 99

	// Now let's break it down.
	// Step 1: Create the orginal slice
	// orginal := []int{10, 20, 30}

	// Stpe 2: Create a destination slice
	// copied := make([]int, len(orginal))
	// This creates a new slice with three elements and its initial values are: [0 0 0]

	// Step 3: Copy the elements
	// copy(copied, orginal)
	// Now the destination contains:
	// [10 20 30]

	// Step 4: Modify the copied slice
	// copied[0] = 99

	fmt.Println("Orginal:", orginal) // Output: Original: [10 20 30]
	fmt.Println("Copied:", copied)   // Output: Copied:   [99 20 30]

	// Important: The orginal remains unchanged because the destination has its own underlyingarray.
}

// 9.3 How does copy() actually work?
// Let's visualize the copying process.

// Before copy()
// Source: orginal
// 10		20		30
// 		copy(copied, orginal)
// Destination: copied
// 10		20		30
// Each element's value is copied into a separate destination element.

// Think of it like this:
// . Source slice contains the orginal data.
// . Destination slice provides seperate storage.
// . copy() transfers element values into that storage.

// The two slice now have independent elements storage.

// 9.4 Very important: Destination length matters
// This is one of the most common mistakes beginners make.
// The copy() function copies only as many elements as the destination can hold.

// The actual number copied is:
//								copied elements = min(len(dst), len(src))

// Let's understand with examples:
// Example A: Same length
func sameLength() {
	source := []int{10, 20, 30}

	destination := make([]int, 3)

	count := copy(destination, source)

	fmt.Println(destination) // Output: [10 20 30]
	fmt.Println(count)       // Output: 3
	// There elements were copied.
}

// Example B: Destination is smaller
func destinationSmaller() {
	source := []int{10, 20, 30}

	destination := make([]int, 2)

	count := copy(destination, source)

	fmt.Println(destination) // Output: [10 20]
	fmt.Println(count)       // Output: 2
	// Why only two?
	// Because the destination has room for only the two elements.
}

// Example C: Destination is larger
func destinationLarger() {
	source := []int{10, 20, 30}

	destination := make([]int, 5)

	count := copy(destination, source)

	fmt.Println(destination) // Output: [10 20 30 0 0]
	fmt.Println(count)       // Output: 5
	// Only three elements were copied. The remaning destination elements retain their orginal zero values.

	// Remember: copy() does not automatically increase the destination's length.
}

// 9.5 Why do we use make() before copy()?
// Consider:
func makeBeforeCopy() {
	orginal := []int{10, 20, 30}

	var copied1 []int
	copy(copied1, orginal)

	fmt.Println(copied1) // Ouput: []
	// Why?
	// Because copied is an nil slice with length zero.
	// There is no destination element availabe to receive the copied values.

	// The correct approach:
	copied := make([]int, len(orginal))
	copy(copied, orginal)

	fmt.Println(copied) // OUtput: [10 20 30]
	// Now the destination has enough element's

	// Important distinction:
	// make([]int, 3)	// Length 3
	// make([]int, 0, 3)	// Length 0, capacity 3

	// The second example still cannot received three elements with copy() because its length is zero.
	// You need to either allocate the desired length or extand the slice first.
}

// 9.6 copy() vs append() - understand the difference
// Both can be used to produce independent slice strong, but they work differently.

// Features								copy()								append()
// Main Purpose							Copy elements into destination		Add elements
// Destination							Must have sufficient length			Grow slice through
// Automatically grows destination		No									Yes
// Returns								Number of elements copied			Updated slice
// Common usage							Copying slices						Adding elements

// You can also copy a slice using append():
func appendCopySlice() {
	orginal := []int{10, 20, 30}

	copied := append([]int(nil), orginal...)
	fmt.Println("Copied:", copied) // Output: [10 20 30]

	copied[0] = 99

	fmt.Println(orginal) // Output: [10 20 30]
	fmt.Println(copied)  // Output: [99 20 30]
	// For now, use make() plus copy() when learning because it make the destination and source relationship very clear.
}

// 9.7 Real-world Example: Copying a product list
// Imagine you are building an inventory management system.

// You have a list of products:
// products := []string{
// 	"Laptop",
// 	"Keyboard",
// 	"Mouse"
// }

// You want another list that you can modify without changing the orginal.
func realWorldExample() {
	products := []string{
		"Laptop",
		"Keyboard",
		"Mouse",
	}

	backup := make([]string, len(products))

	copy(backup, products)

	backup[0] = "Mointor"

	fmt.Println("Orginal Product:", products) // Ouput: Orginal Product: [Laptop Keyboard Mouse]
	fmt.Println("Backup Product:", backup)    // Output: Backup Product: [Mointor Keyboard Mouse]
	// This is useful when you  need a separate copy of a list before modifying it.
}

// 9.8 Comman mistakes and best practices
// Mistake											Correct Approach
// Copying into a nil slice							Allocate destination first
// Using destination capacity instead of length		Ensure destination length is sufficient
// Expecting copy() to grow slice					Remember it copies only availabe elements
// Reversing source and destination					Use copy(destination, source)
// Thinking a := b copies elements					Use copy() for separate storage

// One more important details: copy() performs a shallow copy.
// For slice of integer or string, the behavior we've learned is straight forward. But if you slice contains pointers, maps or structs with reference-type
// field, copying the elements does not recursively duplicate the data those references point to.

// You don't need to study deep copying yet, but remember that copy() does not automatically duplicate everything inside nested data structures.

// What to Remember
// There are only three essential ideas:
//	1. copy(destination, source) copies elements from the source into the destination.
//	2. The destination must have enough length to receive the elements.
//	3. If you allocate the destination separately, modifying its copied elements won't affect the orginal elements.

// 10. The underlyingarray - the most important slice trap
// 10.1 What is underlyingarray?
// A slice in Go does not directly store all its elements.
// Instead, a slice is a small strcutre that refers to an array in memor.

// Think of it like this:
//	. underlyingarray: The actual storage containing the elements.
//	. Slice: A view that let's you access some or all of those elements.

// For example:
// numbers := []int{10, 20, 30, 40, 50}

// Conceptually, Go creates something like this:
// Underlying array
// 		10		20		30		40		50
// 		[0]		[1]		[2]		[3]		[4]
// numbers slice refers to this array

// The slice has three important properties:
// len(numbers)	// Output: 5
// cap(numbers)	// Output: 5
// . len: Number of elements in the slice.
// . cap: Number of elements available from the slice's starting position to the end of its underlying array.

// 10. 2 The slice trap: Two slice can share one array!
// Look at this code.
func sliceTrap1() {
	numbers := []int{10, 20, 30, 40, 50}

	part := numbers[1:4]

	// What does numbers[1:4] means?
	// It selects elements starting at index 1 and stoping before index 4.

	// Therefore:
	fmt.Println(numbers) // Output: [10 20 30 40 50]
	fmt.Println(part)    // Output: [20 30 40]

	// But here's the important part:
	// part does not create a new underlying array
	// So both slice refer to the same array.

	// One underlying array
	// 		10		|20		30		40 |		50
	// 		[0]		|[1]	[2]		[3]|		[4]

	// numbes uses all five elements
	// part view index 1, 2 and 3
}

// 10.3 The suprising behavior
// Now chaning an element thorough part:
func suprisingBehavior() {
	numbers := []int{10, 20, 30, 40, 50}

	part := numbers[1:4]

	part[0] = 999

	fmt.Println("Numbers:", numbers)
	fmt.Println("Part:", part)
	// What do you think happens?
	// Actual Ouptput:
	// Numbers: [10 999 30 40 50]
	// Part: [999 30 40]

	// Why?
	// Because part[0] and numbers[1] refers to the same position in the underlying array.
	// Chaing one changes what both slice see.
	// This is the slice trap!.
}

// 10.4 Another important trap: append()
// Consider:
func anotherAppendTrap() {
	numbers := []int{10, 20, 30, 40, 50}

	part := numbers[1:3]

	part = append(part, 999)

	fmt.Println(numbers)
	fmt.Println(part)

	// You might expect numbers to remain unchanged.
	// But the output is:
	// 	[10 20 999 40 50]
	// [20 30 999]

	// Why?
	// Because part had access to unused capacity in the same underlying array. append() could use that existing space rather than allocating a new array.
	// This means append() can somethimes modify the original slice's underlying array.
	// Important: If the slice doesn't have enough capacity, append() allocates a new underlying array  instead.
}

// 10.5 How to avoid this trap
// If you want a completely independent slice, use copy() build-in Go function:
// For example
func avoidTrap() {
	numbers := []int{10, 20, 30, 40, 50}

	part := numbers[1:4]

	independent := make([]int, len(part))
	copy(independent, part)

	independent[0] = 99

	fmt.Println(numbers)     // Ouput: [10 20 30 40 50]
	fmt.Println(independent) // Output: [99 30 40]
	// Now the two slice have separate underlying arrays.
}

// 11. Passing slice to functions
// 11.1 How to pass a slice to a function
// // You already know how to create a function:
func greet() {
	fmt.Println("Hello!")
}

// You can also pass a slice as a function argurement.
// For Example:
func printNumnbers(numbers []int) {
	fmt.Println("Passing a slice as a function argument:", numbers) // Output: [10 20 30 40 50]
}

// Let's understand:
// func printNumbers(numbers []int)
// This means:
//	. printNumbers is function name.
//	. numbers []int means the function accepts a slice of integers.
//	. When you call printNumbers(numbers) from main function, you pass your slice to the function name called printNumbers.

// Important: You don't need to specify the length of a slice in a function	paramenter.
// This is valid:
// func printNumbers(numbers []int)

// This is not valid Go slice parameter syntax:
// func printNumbers(numbers [4]int)

// The second one accepts an array of exactly four integers, not slice.

// 11.2 Can a function modify the orginal slice?
// Yes, a function can modify the orginal slice's elements.
// Look at this:
func modifyNumbers(numbers []int) {
	numbers[0] = 999 // Output: [99, 20, 30, 40]
	// Why did the orginal slice change?
	// Remember our previours topic: underlying arrays
	// When you pass a slice to a function, Go copies the slice value(its slice descriptor), but the copied descriptor still refers to the same underlying array.
}

// 11.3 Can a function append to a slice?
// Yes, but there's an important details.
// Consider:
func addNumbers(numbers []int) {
	numbers = append(numbers, 50)
	fmt.Println("Inside function:", numbers)
	// Why?
	// Because the function recevies a copy of the slice descriptor.
	// When you do:
	// numbers = append(numbers, 50)

	// The function updates its own local slice descriptor. The caller's slice descriptor reamins unchanged.
	// Because even if append() reuse the orginal underlying arry, the caller's length is still 4.

	// Remember:
	//	. Modifying existing elements can affect the orginal slice.
	//	. Appending inside a function does not automatically update the caller's slice length.
}

// 11.4 How to update the original slice using append()
// Return the updated slice
func addNumber(numbers []int) []int {
	numbers = append(numbers, 50)
	return numbers
}

// 12. Taking a user-input in slice
// 12.1 Taking user input for a fix numbers of elements
// Suppose you want the user to enter 5 numbers.
// You can create an empty slice and use append() to add each input.
func userInput() {
	numbers := []int{}

	var num int

	for i := 0; i < 5; i++ {
		fmt.Print("Enter a number:")
		fmt.Scan(&num)

		numbers = append(numbers, num)
	}

	fmt.Println("Your slice:", numbers)
	// Example input:
	// 	Enter a number: 10
	// Enter a number: 20
	// Enter a number: 30
	// Enter a number: 40
	// Enter a number: 50

	// Output:
	// Your slice: [10, 20, 30, 40, 50]

	// Understand the important parts
	// 1. Create an empty slice
	// numbers := []int{}

	// 2. Declare a variable to store user input
	// var num int

	// 3. Take input
	// fmt.Scan(&num)
	// The & give Scan() the address where it should store the entered value.

	// 4. Add the value to the slice.
	// numbers := append(numbers,  num)
	// Each time the loop runs, append numbers is added in the numbers slice.
}

// 12.2 Let the user decide how many numbers to enter
// This is more flexible.
// Instand of deciding that the user must enter 5 numbers, ask them how many numbers they want to enter.
func askUser() {
	var size int

	fmt.Print("How many numbers do you want to enter?")
	fmt.Scan(&size)

	numbers := []int{}

	var num int

	for i := 0; i < size; i++ {
		fmt.Printf("Enter a number %d:", i+1)
		fmt.Scan(&num)

		numbers = append(numbers, num)
	}
	fmt.Println("Your slice:", numbers)
	// Example input and output:
	// 	How many numbers do you want to enter? 3
	// Enter number 1: 100
	// Enter number 2: 200
	// Enter number 3: 300

	// Your slice: [100 200 300]

	// Notice that size controls how many times the loop executes.
}

// 12.3 Using make() instend of an empty slice
// You can also create a slice with a knwo length
func userInputMake() {
	var size int

	fmt.Print("Enter a slice size:")
	fmt.Scan(&size)

	numbers := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Printf("Enter number %d:", i+1)
		fmt.Scan(&numbers[i])
	}

	fmt.Println("Your slice:", numbers)
	// Example input:
	// 	Enter slice size: 4
	// Enter number 1: 5
	// Enter number 2: 10
	// Enter number 3: 15
	// Enter number 4: 20

	// Output: [5 10 15 20]

	// What's the difference between these two approaches?
	// Empty slice + append						make + indexing
	// numbers := []int{}						numbers := make([]int, size)
	// Use append()								Use numbers[i]
	// Length grows as you append				Length is already set
	// Good when building a slice gradually		Good when size is known

	// Both are correct. Lern to use both.
}

func main() {
	storeMarks()
	arraySlice()
	sliceWithValues()
	declareNilSlice()
	emptySlice()
	makeFunction()
	lenCap()
	accessingUpdatingSliceElements()
	loopWithRange()
	onlyValue()
	appendExample()
	appendExample2()
	appendMultiple()
	nilSlice()
	appendEmptySlice()
	appendLengthCapacity()
	appendOneSliceToAnother()
	modifyAppend()
	slicingExample()
	commanSlicingPatterns()
	selectEntireSlice()
	copyExample()
	copySliceExample()
	sameLength()
	destinationSmaller()
	destinationLarger()
	makeBeforeCopy()
	appendCopySlice()
	realWorldExample()
	suprisingBehavior()
	anotherAppendTrap()
	avoidTrap()

	// 11.1
	numbers := []int{10, 20, 30, 40}
	printNumnbers(numbers)

	modifyNumbers(numbers)
	fmt.Println(numbers)

	addNumbers(numbers)
	fmt.Println("Outside function:", numbers)

	numbers = addNumber(numbers)
	fmt.Println("Main function:", numbers)
	// Notice:
	// numbers = addNumber(numbers)

	// We assign the returned slice back to orignal numbers.
	// This is important because append() might allocated a new underlaying array, and the returned slice conatins the correct updated descriptor.

	userInput()
	askUser()
	userInputMake()
}
