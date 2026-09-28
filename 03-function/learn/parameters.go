package main

import "fmt"

// 1. What is Parameters?
// A parameters is a variable that receives a value when a function is called.

// 2. Why do we need parameters?
// Suppose you want to greet different people.
// Without paramenters, you might write a seprate functions.

func greetBishan() {
	fmt.Println("Hello, Bishan")
}

func greetRam() {
	fmt.Println("Hello, Ram")
}

// This is repetitive.
// Instance, create one function that accepts a name:

func greet(name string) {
	fmt.Println("Hello", name)
}

// Now:
// greet("Bishan")
// greet("Ram")
// greet("Sita")

// Output:
// Hello Bishan
// Hello Ram
// Hello Sita

// The function is reuseable because the input can change.

// 3. Parameter vs Argument
// This is one of the most important things to understand.

// Look at:
func greet(name string) {
	fmt.Println("Hello", name)
}

// Here:
name string // is the parameters.

// When we call:
greet(Bishan) // "Bishan" is the argument.

// Simple distinction
// Paramenter -> variable defined in the function
// Argument -> actual value passed to the function

// Example:
func greet(name string) {	// name = parameters
	fmt.Println("Hello", name)
}

greet("Bishan")	// "Bishan" = argument

// Think:
// Function definition:
//         parameter
//            ↓
// func greet(name string)

// Function call:
//            argument
//               ↓
//        greet("Bishan")

// 4. How a paramenter works
// Consider:

func greet(name string) {
	fmt.Println("Hello", name)
}

func main() {
	greet("Bishan")
}

// When Go Sees:
greet("Bishan")

// You can mentally think:
// "Bishan"
//     ↓
//  name
//     ↓
// fmt.Println("Hello", name)

// So inside the function:
name == "Bishan"

// Therefore
// Hello Bishan

// 5. Parameter syntax
// The basic syntax is:

func functionName(parameterName paramterTypes) {
	// code
}

// Example:
func greet(name string) {
    fmt.Println("Hello", name)
}

// Break it down:
// func
//  ↓
// greet
//  ↓
// (name string)
//      ↓
//  parameter

// More specifically:
name string

// Means:
// name -> parameter name
// string -> parameter type

// 6. Parameters behave like local variables
// This is an important mental model.

// When you write.
func greet(name string) {
    fmt.Println(name)
}

// name behaves like a variable that exists inside the function.

// For example:
func greet(name string) {
	message := "Hello" + name

	fmt.Println(message)
}

// Here you have:
// name
//  ↓
// parameter

// message
//  ↓
// local variable

// Both are available inside the function.

// 7. Multiple parameters
// A function can accept multiple parameters.

// Example:
func add(a int, b int) {
	fmt.Println(a + b)
}

// Call it:
add(10, 20)

// Output:
30

// Think:
// add(10, 20)

// 10 → a
// 20 → b

// Inside the function:
a = 10
b = 20

// 8. Multiple paramenters of the same type
// Go Provides a shorter syntax.

// Instead of:
func sub(c int, d int) {
	fmt.Println(a - b)
}

// You can write:
func sub(c, d int) {
	fmt.Println(a-b)
}

// Both mean the same thing.
c -> int
d -> int

// This is very common in Go.

// 9. Parameters can have different types
// For example:
func introduce(name string, age int) {
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}

// Call:
introduce("Bishan", 22)

// Output:
Name: Bishan
Age : 22

// This mapping is positional:
// "Bishan" → name
// 22       → age

// So this:
introduce("Bishan", 22)

// matches:
func introduce(name string, age int)

// 10. Order Matters
// Parameters are matched by position.

// Example:
func introduction(name string, age int) {
	fmt.Println(name, age)
}

// Correct
introduction("Bishan", 22)

// Incorrect
introduction(22, "Bishan")

// Why?
// Because Go expects:

// 1st argument → string
// 2nd argument → int

// but you gave:
// 1st argument → int
// 2nd argument → string

// Go's type system catches this error.

// 11. Parameters make functions reusable
// This is the real reason parameters matter.

// Without parameters:
func squareOfFive() {
	fmt.Println(5 * 5)
}

// That's useful only for 5

// With a parameter:
func square(number int) {
	fmt.Println(number * number)
}

// Now:
square(5)
square(10)
square(20)

// Output:
25
100
400

// One function handles many inputs.

// 12. Parameters + if/else
// This connects directly to what you've already learned.

func checkAge(age int) {
    if age >= 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }
}

// Call:
checkAge(22)
checkAge(15)

// Output:
Adult
Minor

// The flow is:
// checkAge(22)
//      ↓
// age = 22
//      ↓
// if age >= 18
//      ↓
// Adult

// This is where functions start becoming genuinely useful.

// 13. Parameters do not have to be named value
// You choose the parameter name.
// These are equivalent:
func greet(name string) {
    fmt.Println("Hello", name)
}

// and:
func greet(person string) {
    fmt.Println("Hello", person)
}

// add:
func greet(x string) {
    fmt.Println("Hello", x)
}

// But prefer meaningful names:
func greet(name string)

// rather than:
func greet(x string)

// Good parameter names make your code easier to understand.

// 14. A parameter only exists inside its function
// Consider:
func greet(name string) {
    fmt.Println(name)
}

func main() {
    greet("Bishan")

    fmt.Println(name) // ERROR
}

// Why?

// Because name belongs to greet.

// Think of it as:
// greet()
// ┌───────────────────┐
// │ name              │
// │                   │
// │ only exists here  │
// └───────────────────┘

// main()
// ┌───────────────────┐
// │ name ❌            │
// └───────────────────┘

// This is related to scope, which you already learned with variables.

// 15. Parameters can receive expressions
// The argument doesn't always have to be a literal value.
// You can pass a variable:

age := 22

checkAge(age)

// You can also pass an expression:
checkAge(10 + 12)

// Go evaluates:
// 10 + 12
//  ↓
// 22
//  ↓
// checkAge(22)

// You can even do:
checkAge(20 - 5)

// which becomes:
checkAge(15)