package main

// What is a return value?
// A parameter sends data into a function. A return value sends data back out of a function.

// 1. Why do we need return values?
// Return values are useful because they allow us to get data back from a function.
// This is important because it allows us to use the result of a function in other parts of our code.

// 2. How do we use return values?
// We can use return values by assigning the result of a function call to a variable.
// For example, if we have a function that adds two numbers together, we can assign the result of that function to a variable and then use that variable in other parts of our code.

// 3. How do we define return values?
// We define return values by specifying the type of data that the function will return in the function signature.
// For example, if we want our function to return an integer, we would specify "int" as the return type in the function signature.

// 4. How do we return values from a function?
// We return values from a function using the "return" keyword followed by the value we want to return.
// For example, if we want to return the sum of two numbers, we would use "return a + b" where "a" and "b" are the two numbers being added together.

import "fmt"

func main() {
	// Call the add function and assign the result to a variable
	result := add(3, 5)
	// Print the result
	fmt.Println("The sum is:", result)
}

// add is a function that takes two integers as parameters and returns their sum as an integer
func add(a int, b int) int {
	// Return the sum of a and b
	return a + b
}
