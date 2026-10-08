package main

import "fmt"

// 1. Return a greeting: Write greeting() string to return "Welcome to Go!". Store the result in a variable and print it.
func greeting() string {
	return "Welcome to Go!"
}

// 2. Add two numbers: Write add(a, b int) int to return their sum.
func add(a, b int) int {
	sum := a + b
	return sum
}

// 3. Square a number: Write square(n int) int. Test it with a positive number, a negative number, and zero.
func square(n int) int {
	return n * n
}

// 4. Calculate rectangle area: Write rectangleArea(length, width float64) float64.
func rectangleArea(length, width float64) float64 {
	return length * width
}

// 5. Check even numbers: Write isEven(n int) bool. Use the returned result in an if/else statement in main().
func isEven(n int) bool {
	return n%2 == 0
}

// 6. Return the larger number: Write max(a, b int) int. Test equal numbers too.
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 7. Find an absolute value: Write absolute(n int) int to return a nonnegative value. For example, absolute(-8) should return 8.
func absolute(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

// 8. Convert temperature: Write celsiusToFahrenheit(celsius float64) float64. Use:
// Fahrenheit = Celsius × 9 / 5 + 32
// Check that 0°C produces 32°F.
func celsiusToFahrenheit(celsius float64) float64 {
	fahrenheit := celsius*9/5 + 32
	return fahrenheit
}

// 9. Calculate a sum using a loop: Write sumToN(n int) int to return the sum from 1 through n. Return 0 when n <= 0. For example, sumToN(5) should return 15.
func sumToN(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	return sum
}

// 10. Return a grade: Write grade(score int) string with these rules:
// Score			Return
// Outside 0–100	"Invalid"
// 90–100			"A"
// 80–89			"B"
// 70–79			"C"
// 60–69			"D"
// 0–59				"F"

// Test the boundaries, including 59, 60, 89, and 90.

func grade(score int) string {
	var grade string
	if score < 0 || score > 100 {
		grade = "Invalid"
	} else if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else if score >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}

	return grade

}

func main() {
	// 1.
	result := greeting()
	fmt.Println(result)

	// 2.
	fmt.Println(add(3, 4))

	// 3.
	fmt.Println(square(7))
	fmt.Println(square(-7))
	fmt.Println(square(0))

	// 4.
	fmt.Println(rectangleArea(10, 5))

	// 5.
	if isEven(8) {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 6.
	results := max(10, 20)
	fmt.Println(results)

	// Equal numbers
	results2 := max(7, 7)
	fmt.Println(results2)

	// 7.
	fmt.Println(absolute(-7))
	fmt.Println(absolute(10))
	fmt.Println(absolute(0))

	// 8.
	fmt.Println(celsiusToFahrenheit(34.8))
	fmt.Println(celsiusToFahrenheit(0))

	// 9.
	fmt.Println(sumToN(5))

	// 10.
	fmt.Println(grade(59))
	fmt.Println(grade(69))
	fmt.Println(grade(79))
	fmt.Println(grade(89))
	fmt.Println(grade(90))
	fmt.Println(grade(101))
	fmt.Println(grade(0))

}
