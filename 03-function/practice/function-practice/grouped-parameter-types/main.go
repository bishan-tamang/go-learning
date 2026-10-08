package main

import "fmt"

// 1. Add two integers: Write printSum(a, b int) to print their sum. Call it with two different pairs.
func printSum(a, b int) {
	fmt.Println(a + b)
}

// 2. Multiply three integers: Write printProduct(a, b, c int) to print their product.
func printProduct(a, b, c int) {
	fmt.Println(a * b * c)
}

// 3. Greet a person: Write greet(firstName, lastName string) to print a greeting using their full name.
func greet(firstName, lastName string) {
	fmt.Println("Hello,", firstName, lastName)
}

// 4. Rectangle measurements: Write printRectangle(length, width float64) to print its area and perimeter.
func printRectangle(length, width float64) {
	fmt.Println("Area:", length*width)
	fmt.Println("Perimeter:", 2*(length+width))
}

// 5. Find the larger number: Write printLarger(a, b int) to print the larger value. If they are equal, print "Equal".
func printLarger(a, b int) {
	if a > b {
		fmt.Println(a)
	} else if b > a {
		fmt.Println(b)
	} else {
		fmt.Println("Equal")
	}
}

// 6. Find the smallest number: Write printSmallest(a, b, c int) to print the smallest of three integers.
func printSmallest(a, b, c int) {
	if a <= b && a <= c {
		fmt.Println("smallest:", a)
	} else if b <= a && b <= c {
		fmt.Println("smallest:", b)
	} else {
		fmt.Println("smallest:", c)
	}
}

// 7. Print a range: Write printRange(start, end int) to print every integer from start to end, inclusive. Assume start <= end.
func printRange(start, end int) {
	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
}

// 8. Use two parameter groups: Write showProfile(name, country string, age, score int) to print a person's details. Identify the two groups sharing types.
func showProfile(name, country string, age, score int) {
	fmt.Println(name, country, age, score)
}

// name, country string // Both are string
// age, score int       // Both are int

// 9. Combine grouped and separate parameters: Write printCalculation(a, b int, operation string).
// Use switch to print the result for "add", "subtract", or "multiply".
func printCalculation(a, b int, operation string) {
	switch operation {
	case "add", "+":
		fmt.Println("Add:", a+b)
	case "subtract", "-":
		fmt.Println("Subtract:", a-b)
	case "multiply", "*":
		fmt.Println("Multiply:", a*b)
	case "divide", "/":
		if b == 0 {
			fmt.Println("Can't Divided by zero.")
		} else {
			fmt.Println("Division:", a/b)
		}
	default:
		fmt.Println("Invalid Operation")
	}
}

// 10. Rewrite ungrouped parameters: Rewrite this signature using grouped types, then implement it to print all six values:
// func showDetails(firstName string, lastName string, age int, score int, height float64, weight float64)
func showDetails(firstName, lastName string, age, score int, height, weight float64) {
	fmt.Println(firstName, lastName, age, score, height, weight)
}

func main() {
	// 1.
	printSum(3, 4)
	printSum(7, 7)

	// 2.
	printProduct(2, 2, 2)

	// 3.
	greet("Bishan", "Tamang")

	// 4.
	printRectangle(10, 5)

	// 5.
	printLarger(7, 5)
	printLarger(5, 7)
	printLarger(7, 7)

	// 6.
	printSmallest(7, 7, 9)
	printSmallest(9, 7, 7)
	printSmallest(7, 9, 7)

	// 7.
	printRange(1, 10)

	// 8.
	showProfile("Bishan", "Nepal", 22, 99)

	// 9.
	printCalculation(10, 2, "add")
	printCalculation(10, 2, "subtract")
	printCalculation(10, 2, "multiply")
	printCalculation(10, 2, "divide")
	printCalculation(10, 0, "/")
	printCalculation(10, 2, "%")

	// 10.
	showDetails("Bishan", "Tamang", 22, 100, 5.7, 42.450)
}
