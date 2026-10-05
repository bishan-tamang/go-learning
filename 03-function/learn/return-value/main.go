package main

import "fmt"

// What is return value of the function?
// A return value is a result that a function sends back to the code that called it.
// for example, a function can receive two numbers, add them, and return the answer:

func add(a, b int) int {
	return a + b
}

// Let's build your understanding step by step.

// 1. Understand input, work, and output
// Think of a function like a calculator:

//   . Parameters are its input.
//   . Function body does the work.
//   . Return value is its output.

func sum(a, b int) int {
	return a + b
}

// Here is what each part means:
// Part					Meaning
// func			        Defines a function
// add					Function name
// a, b int				Two input parameters
// int after )			Types of value the function returns
// return a + b		    Sends the result back to the caller

// When you call:
// sum(3, 4)
// The function receives 3 and 4 as input, adds them together, and returns 7 as output.

// 2. Store and use a returned value
// Here is a complete program:

func addition(a, b int) int {
	return a + b
}

// func main() {
// 	addtionResult := addition(3, 4)
// 	fmt.Println(addtionResult)
// }

// Output: 7

// Follow the execution:
// 1. main() calls addition(3, 4)
// 2. Inside addition, a become 3 and b becomes 4
// 3. return a +b sends 7 back
// 4. addtionResult receives the return value, which is 7
// 5. fmt.Println(additionresutl)prints it.

// You can also use the return value directly without storing it in a variable:
// fmt.Println(addition(3,4)) // Output: 7

// or use it in another calculation:
// total := addition(3, 4) * 2
// fmt.Println(total)	// Output: 14

// The function call produces a value that you can store, print, compare or use in another expression.
// You can use the return value in any way you would use a value of that type.

// 3. Understand returning versus printing
// These functions do different jobs:

func printSum(a, b int) {
	fmt.Println(a + b)
}

func returnSum(a, b int) int {
	return a + b
}

// Printing											returning
// Displays somethings in the terminal.				Gives a result to the caller.
// Use fmt.Println(...)								Uses return ...
// Does not itselft provide a reusable resutl.		Lets other code use the result.

// Usages:
// func main() {
// 	printSum(3, 4) // Output: 7

// 	result := returnSum(3, 4)
// 	fmt.Println(result) // Output: 7

// 	// This is invalid:
// 	// result2 := printSum(3, 4)
// 	// Why? printSum does not return a value, so you cannot assign its result to a variable.

// 	// For calculations and decisions, returning a result usually makes the functiion more easier to reuse.

// }

// 4. Return different types
// Functions can return different types, but the returned value must match the declared return type.

// An integer:
func square(x, y int) int {
	return x * y
}

// A string:
func greet(name string) string {
	return "Hello, " + name
}

// A boolean:
func isEven(n int) bool {
	return n%2 == 0
}

// A float
func divide(a, b float64) float64 {
	return a / b
}

// Usages:
// func main() {
// 	// integer
// 	squareResult := square(3, 4)
// 	fmt.Println(squareResult) // Output: 12

// 	// string
// 	greeting := greet("Mamita")
// 	fmt.Println(greeting) // Output: Hello, Mamita

// 	// boolean
// 	even := isEven(4)
// 	fmt.Println(even) // Output: true

// 	// float
// 	divisionResult := divide(10.0, 2.0)
// 	fmt.Println(divisionResult) // Output: 5
// }

// A function can also return a value without receving parameters.
func defaultLanguage() string {
	return "Go"
}

// Usage:
// func main() {
// 	language := defaultLanguage()
// 	fmt.Println(language) // Output: Go
// }

// Parameters and return values are independent: a function can have either, both or neither.
// A function can receive parameters, return a value, do both, or do neither.

// 5. return immediately ends the current function call.
// Onec a return runs, the function stops executing and returns the value to the caller.
func checkAge(age int) string {
	if age < 18 {
		return "You are a minor"
	}
	return "You are an adult"
}

// For checkAge(15):
//	. The condtion is true.
// 	. The function returns "You are a minor".
// 	. The function ends, and the second return statement is never reached.

// For checkAge(20):
//	. The condition is false.
// 	. The first return statement is skipped.
// 	. Execution continues.
// 	. The function returns "You are an adult".

// This is called "early return" and is a common pattern in Go.
// It can make your code easier to read by handling special cases first and keeping the main logic at the bottom of the function.

func discountPrice(price float64, discount float64) float64 {
	if discount <= 0 {
		return price // No discount, return original price
	}
	return price - price*discount/100 // Apply discount and return new price
}

// func main() {
// 	fmt.Println(discountPrice(1000, 10)) // Output: 900
// 	fmt.Println(discountPrice(1000, 0))  // Output: 1000
// }

// In a function with no return values, a bare return simply exits:
func greetUser(name string) {
	if name == "" {
		fmt.Println("Hello, Guest!")
		return // Exit the function early if no name is provided
	}
	fmt.Println("Hello, " + name + "!")
}

// Usage:
// func main() {
// 	greetUser("Alice") // Output: Hello, Alice!
// 	greetUser("")      // Output: Hello, Guest!
// }

// 6. Make sure evey possible path returns
// If a function promises a result, it cannot exit without returning a value. The Go compiler will check this for you and give an error if a return is missing.

// Incorrect:
// func ageCheck(age int) string {
// 	if age >= 18 {
// 		return "You are an adult"
// 	}
// }

// What should the function return if age is 15? There is no answer, So Go reports a missing return.

// Correct:
func ageCheckCorrect(age int) string {
	if age >= 18 {
		return "You are an adult"
	}
	return "You are a minor"
}

// The function now returns a value for every possible path through the code.

// Also, the type must match. If a function is declared to return a string, it cannot return an integer or any other type.
// The Go compiler will check this for you and give an error if the types do not match.
// func getAge() int {
// 	return "25" // Error: cannot use "25" (type string) as type int in return statement
// }

// Correct:
func getAgeCorrect() int {
	return 25 // Correct: returns an integer
}

// 7. Return multiple values
// Go functions can return more than one value.
// Declare the return types inside parentheses:
func rectangle(length int, width int) (int, int) {
	area := length * width
	perimeter := 2 * (length + width)

	return area, perimeter
}

// Receive both results in the same order:
// func main() {
// 	area, perimeter := rectangle(5, 3)
// 	fmt.Println("Area:", area)         // Output: Area: 15
// 	fmt.Println("Perimeter:", perimeter) // Output: Perimeter: 16
// }

// The order matters:
// First return value -> first variable
// Second return value -> second variable

// Return types can also different types:
func getUserInfo() (string, int) {
	name := "Alice"
	age := 30
	return name, age
}

// Usage:
// func main() {
// 	name, age := getUserInfo()
// 	fmt.Println("Name:", name) // Output: Name: Alice
// 	fmt.Println("Age:", age)   // Output: Age: 30
// }

// If you only need one result then you can ignore the other by using an underscore (_):
// func main() {
// 	name, _ := getUserInfo()
// 	fmt.Println("Name:", name) // Output: Name: Alice
// }

// 9. Return multiple values from a function that takes no parameters:
func getCoordinates() (float64, float64) {
	x := 10.5
	y := 20.3
	return x, y
}

// Usage:
// func main() {
// 	x, y := getCoordinates()
// 	fmt.Println("X:", x) // Output: X: 10.5
// 	fmt.Println("Y:", y) // Output: Y: 20.3
// }

// 8. Return a result and an error
// In Go, it is common to return a result along with an error value to indicate if something went wrong.
func divideWithError(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Usage:
func main() {
	result, err := divideWithError(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result) // Output: Result: 5
	}

	result, err = divideWithError(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: cannot divide by zero
	} else {
		fmt.Println("Result:", result)
	}
}

// This pattern allows the caller to handle errors gracefully and is widely used in Go programming.

// 9. Named return values
// You can name the return values in the function signature:
func calculate(length int, width int) (area int, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // Return the named values
}

// Usage:
// func main() {
// 	area, perimeter := calculate(5, 3)
// 	fmt.Println("Area:", area)         // Output: Area: 15
// 	fmt.Println("Perimeter:", perimeter) // Output: Perimeter: 16
// }

// Named return values can make your code more readable, especially for functions with multiple return values.
