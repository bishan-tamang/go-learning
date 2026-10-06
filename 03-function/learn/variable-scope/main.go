package main

import "fmt"

// Variable scope means where a variable's name can be used in your code.

// Start with three places you can declare varibales:
// Scope Level		Where it is declared								Visibility
// Function Scope	Inside a function body								Accessible anywhere inside the function, including nested blocks.
// Package Scope	Outside any function (starts with lowercase)		Accessible by all files within the same package directory.
// Block Scope		Inside curly braces {} (e.g., if, for, switch)		Only accessible within that specific code block.

// let's learn them step by step

// 1. Variable inside a function
// A variable declared inside a function is a local variable.

func greet() {
	name := "Bishan"
	fmt.Println(name) // Works
}

// You cannot directly use that variable in another function:

func main() {
	greet()
	fmt.Println(name)
}

// Calling greet() does not make its local variable available in main.

// 2. The same name in different functions.
// Two functions can have separate variables with the same name:

func showAge() {
	age := 22
	fmt.Println("Inside showAge:", age)
}

func main() {
	age := 30

	showAge()
	fmt.Println("Inside main:", age)
}

// Output:
// Inside showAge: 22
// Inside main: 30

// These are two different variables. Creating age inside showAge does not change the age inside main.

// 3. Parameters are local to theri function.
// Parameters can be used inside the function that declares them:
func double(number int) int {
	return number * 2 // number is available here
}

// But they aren't available in main:
func main() {
	result := double(7)
	fmt.Println(result) // Works
	// fmt.Println(number) // Error: undefined: number
}

// 4. Variable inside an if block
// A block is group of statements enclosed by {}.
func main() {
	age := 22

	if age >= 18 {
		message := "You are an adult"

		fmt.Println(age)     // Works
		fmt.Println(message) // Works
	}
	fmt.Println(age) // Works
	// fmt.Println(message) // Error: undefined: message
}

// Here:
// age belongs to the surrounding function block, so the inner if block can use it.
// message belongs to the if block, so code outside that block cannot use it.

// An inner block can use variable from its enclosing block. The enclosing block cannot use variables declared only inside the inner block.

// If you need message afterward, declare it before the if.
func main() {
	age := 22
	message := "You are under 18"

	if age >= 18 {
		message = "You are an adult"
	}

	fmt.Println(message)
}

// Notic message = ....: it updates the existing variable.

// 5. Shadowing: same name in an inner block
// An inner block can declare a new variable with the same name as an outer variable. This is called shadowing.

func main() {
	age := 22

	if age >= 18 {
		age := 30
		fmt.Println("Inside if:", age)
	}

	fmt.Println("Outside if:", age)
}

// Output:
// Inside if: 30
// Outside if: 22

// Inside the if, age := 30 create a new variable that temporily hides the outer age.

// Compare that with assignment
func main() {
	age := 22

	if true {
		age = 30
	}

	fmt.Println(age) // Output: 30
}

// Here, age = 30 updates the existing outer variable.

// Inside the inner block		Meaning
// age := 30					Create a new local variable; shadow the outer one
// age = 30						Update the eixiting outer variable

// Avoid accidental shadowing when you intend to update a variable.

// 6. Package-level variables
// A variable declared outside all function is a package-level variable
var language = "Go"

func showLanguage() {
	fmt.Println(language)
}

func main() {
	showLanguage()
	fmt.Println(language)
}

// Both functions can use language becasue they belong to the same package.
// At package level, use var; := is allowed only inside functions.
