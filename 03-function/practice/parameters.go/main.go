package main

import "fmt"

// Basic Parameters
// Q1. Greeting
func greet(name string) {
	fmt.Println("Hello,", name)
}

// Q2. Age

func showAge(age int) {
	fmt.Println("Your age is", age)
}

// Q3. Number
func showNumber(num int) {
	fmt.Println("The Number is", num)
}

// Multiple Parameters
// Q4. Addition
func add(a, b int) {
	fmt.Println(a + b)
}

// Q5. Student Information
func studentInfo(name string, age int) {
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}

// Q6. Area of Rectangle
func rectangle(length, width int) {
	fmt.Println("Area of Rectangle is:", length*width)
}

func main() {
	// Greeting
	greet("Bishan")
	greet("Mamita")

	// Show Age
	showAge(22)

	// show Number
	showNumber(7)
	showNumber(10)
	showNumber(11)

	// Addition
	add(7, 7)
	add(45, 55)
	add(10, 10)

	// Student Information
	studentInfo("Bishan", 22)
	studentInfo("Mamita", 21)

	// Area of Rectangle
	rectangle(10, 5)
	rectangle(20, 10)
}
