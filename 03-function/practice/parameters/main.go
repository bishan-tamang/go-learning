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

// Parameters + Operators
// Q7. Square
func square(num int) {
	fmt.Println(num * num)
}

// Q8. Calculator
func calculator(a, b int) {
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
}

// Parameter + if/else
// Q9. Age Checker
func checkAge(age int) {
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
}

// Q10. Even or Odd
func checkNumber(num int) {
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

// Q11. Greater Number
func compareNumber(a, b int) {
	if a > b {
		fmt.Println("A is greater")
	} else if b > a {
		fmt.Println("B is greater")
	} else {
		fmt.Println("Both are equal")
	}
}

// Challenge
// Q12. Grade Checker
func gradeChecker(marks int) {
	if marks < 1 || marks > 100 {
		fmt.Println("Invalid Marks")
	} else if marks >= 90 {
		fmt.Println("Grade: A")
	} else if marks >= 80 {
		fmt.Println("Grade: B")
	} else if marks >= 70 {
		fmt.Println("Grade: C")
	} else if marks >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}
}

// Q13. Temperature Checker
func temperatureChecker(temp float64) {
	if temp >= 35 {
		fmt.Println("Very Hot")
	} else if temp >= 25 {
		fmt.Println("Warm")
	} else if temp >= 15 {
		fmt.Println("Normal")
	} else {
		fmt.Println("Cold")
	}
}

// Q14. Mini ATM
func checkBalance(balance, withdraw float64) {
	if withdraw > 0 && withdraw <= balance {
		fmt.Println("Remaining balance:", balance-withdraw)
	} else {
		fmt.Println("Insufficient balance")
	}
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

	// Square
	square(5)
	square(7)

	// Calculator
	calculator(10, 5)

	// Check Age
	checkAge(16)
	checkAge(22)

	// Check Even or Odd
	checkNumber(25)
	checkNumber(22)

	// Check Greater Number
	compareNumber(10, 5)
	compareNumber(3, 8)
	compareNumber(7, 7)

	// Grade Checker
	gradeChecker(0)
	gradeChecker(101)
	gradeChecker(75)
	gradeChecker(59)

	// Temperature Checker
	temperatureChecker(35.6)
	temperatureChecker(27.50)
	temperatureChecker(19.10)
	temperatureChecker(14.7)

	// Mini ATM
	checkBalance(1000, 300)
	checkBalance(1000, 1100)
}
