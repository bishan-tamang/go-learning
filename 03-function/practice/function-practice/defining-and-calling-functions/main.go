package main

import "fmt"

// 1. Say hello: Define sayHello() to print "Hello, Go!". Call it from main().
func sayHello() {
	fmt.Println("Hello, Go!")
}

// 2. Introduce yourself: Define introduce() to print your name, age, and country on separate lines.
func introduce() {
	fmt.Println("Bishan Tamang")
	fmt.Println(22)
	fmt.Println("Nepal")
}

// 3. Call repeatedly: Define showMessage() to print "Practice makes progress.". Call it three times from main().
func showMessage() {
	fmt.Println("Practice makes progress.")
}

// 4. Print a menu: Define showMenu() to display:
// 1. Add
// 2. Subtract
// 3. Exit
func showMenu() {
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Exit")
}

// 5. Add fixed numbers: Define addNumbers(). Inside it, declare two integers, add them, and print their sum.
func addNumbers() {
	num1 := 3
	num2 := 4
	add := num1 + num2
	fmt.Println(add)
}

// 6. Check a fixed number: Define checkNumber(). Declare an integer inside it and print whether it is positive, negative, or zero.
func checkNumber() {
	num := 7

	if num > 0 {
		fmt.Println("positive")
	} else if num < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("zero")
	}
}

// 7. Count upward: Define countToTen() to print the numbers from 1 to 10 using a loop.
func countToTen() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 8. Print a multiplication table: Define printTable() to print the multiplication table of 7, from 7 × 1 through 7 × 10.
func printTable() {
	for j := 1; j <= 10; j++ {
		result := 7 * j
		fmt.Printf("7 x %d = %d\n", j, result)
	}
}

// 9. Control execution order: Define start(), process(), and finish(). Each prints its own message. Call them in that order from main(), then change the order and observe the output.
func start() {
	fmt.Println("Starting...")
}

func process() {
	fmt.Println("Processing...")
}

func finish() {
	fmt.Println("Finishing...")
}

// 10. Organize a small program: Define showHeader(), showInstructions(), and showFooter(). Call each twice from main() without copying its printing code into main().
func showHeader() {
	fmt.Println("Introduction to Go")
}

func showInstructions() {
	fmt.Println("Read each page.")
}

func showFooter() {
	fmt.Println("Closed the Book")
}

// Calling all function into main function
func main() {
	sayHello()
	introduce()
	showMessage()
	showMessage()
	showMessage()
	showMenu()
	addNumbers()
	checkNumber()
	countToTen()
	printTable()
	// In order
	start()
	process()
	finish()

	// Change the order
	process()
	start()
	finish()

	showHeader()
	showHeader()
	showInstructions()
	showInstructions()
	showFooter()
	showFooter()
}
