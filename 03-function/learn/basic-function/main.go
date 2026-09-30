package main

import "fmt"

// 1. What is a Function?
// A function is a reusable block of code that performs a specific task.

// Instance of repetedly writing:
// fmt.Println("Hello, Bishan")

// You can create a function:
func greetElder() {
	fmt.Println("Hello, Mother")
}

// Then reuse it in main function:
// greetElder()
// greetElder()
// greetElder()

// Think:
// Function
//    ↓
// Input
//    ↓
// Process
//    ↓
// Output

// For Example:
// 10 + 20
//    ↓
// add()
//    ↓
// 30

// 2. Why function matter
// functions helps you:
//	. avoid repeating code
//	. organize large program
//	. make code easier to test
//	. isolate specific responsibilities
//	. reuse logic
//	. make errors easier to handle
//	. build larger programs froms smaller pieces

// A good function usually does one clear job.
// for example:
// calculateTotal()
// is better then one gaint function that calculates totoal, saves to a database, sends an email, and prints a report.

// 3. Function Syntax
// To create a function, do the following:

//	. use the func keyword
//	. Specify a name for the function, followed by parentheses ().
//	. Finally, add code that define what the function should do, inside curly braces {}.
//
// Syntax
func functionName() {
	// code to be executed
}

// For Example,
func myMessage() {
	fmt.Println("I just got executed!")
}

// Breakdown:
// func -> keyword
// myMessage -> function name
// () -> parameters
// {} -> function body

// 4. Calling a function
// Defining a function doesn't execute it.
func greetBrother() {
	fmt.Println("Hi, Bishan")
}

// You need to call it.
// greetBrother()

// Complete Example

// func greet() {
// 	fmt.Println("Hello, Sir")
// }

// func main() {
// 	greet() // call the function in main function
// }

// Output: Hello, Sir
