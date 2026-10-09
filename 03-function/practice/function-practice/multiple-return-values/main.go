package main

import (
	"fmt"
)

// 1. Return personal details: Write getProfile() (string, int) to return your name and age. Store both results and print them.
func getProfile() (string, int) {
	name := "Bishan"
	age := 22

	return name, age
}

// 2. Return sum and difference: Write sumAndDifference(a, b int) (int, int) to return a + b and a - b.
func sumAndDifference(a, b int) (int, int) {
	return a + b, a - b
}

// 3. Return area and perimeter: Write rectangleStats(length, width float64) (float64, float64).
func rectangleStats(length, width float64) (float64, float64) {
	area := length * width
	perimeter := 2 * (length + width)

	return area, perimeter
}

// 4. Swap two values: Write swap(a, b int) (int, int) to return the values in reverse order. Use:
// x, y = swap(x, y)
// Print x and y before and after the call.
func swap(a, b int) (int, int) {
	return b, a
}

// 5. Return quotient and remainder: Write divideAndRemainder(a, b int) (int, int). Assume b is nonzero. For inputs 17 and 5, return 3 and 2.
func divideAndRemainder(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}

// 6. Return minimum and maximum: Write minMax(a, b, c int) (int, int) to return the smallest and largest values. Test a case where all three are equal.
func minMax(a, b, c int) (int, int) {
	minimum := a
	maximum := a

	// (100, 110, 10)
	if b < minimum {
		minimum = b
	}

	if c < minimum {
		minimum = c
	}

	if b > maximum {
		maximum = b
	}

	if c > maximum {
		maximum = c
	}

	return minimum, maximum
}

// 7. Convert total seconds: Write minutesAndSeconds(total int) (int, int). Assume total >= 0. For 135, return 2 minutes and 15 seconds.
func minutesAndSeconds(total int) (int, int) {
	return total / 60, total % 60
}

// 8. Return a value and a status: Write checkEven(n int) (int, bool) to return the original number and whether it is even.
// Use both results in a sentence printed from main().
func checkEven(n int) (int, bool) {
	return n, n%2 == 0
}

// 9. Ignore one result: Write squareAndCube(n int) (int, int). Call it three times:
// - Receive and print both results.
// - Receive only the square using _ for the cube.
// - Receive only the cube using _ for the square.
func squareAndCube(n int) (int, int) {
	square := n * n
	cube := n * n * n

	return square, cube
}

// 10. Return three results: Write numberStats(a, b, c int) (int, int, float64) to return the sum1, product, and average.
// Make sure inputs 1, 2, and 2 produce an average of approximately 1.6667, rather than 1.
func numberStats(a, b, c int) (int, int, float64) {
	sum1 := a + b + c
	product := a * b * c
	average := float64(sum1) / 3

	return sum1, product, average
}

func main() {
	// 1.
	name, age := getProfile()
	fmt.Println(name)
	fmt.Println(age)

	// 2.
	sum, difference := sumAndDifference(10, 5)
	fmt.Println(sum, difference)

	// 3.
	area, perimeter := rectangleStats(14, 7)
	fmt.Println(area, perimeter)

	// 4.
	x, y := 8, 9
	fmt.Println("Before:", x, y)

	x, y = swap(x, y)
	fmt.Println("After:", x, y)

	// 5.
	quotient, remainder := divideAndRemainder(17, 5)
	fmt.Println("Quotient:", quotient, "Remainder:", remainder)

	// 6.
	minimum, maximum := minMax(100, 110, 10)
	fmt.Println("Minimum:", minimum, "Maximum:", maximum)

	minimum, maximum = minMax(7, 7, 7)
	fmt.Println("Minimum:", minimum, "Maximum:", maximum)

	// 7.
	minutes, seconds := minutesAndSeconds(562)
	fmt.Println(minutes, "minutes and", seconds, "seconds")

	// 8.
	value, status := checkEven(87)
	fmt.Println("Orginal number:", value, "Status:", status)

	// 9.
	// Receive and print both results.
	square, cube := squareAndCube(2)
	fmt.Println(square, cube)

	// Receive only the square; ignore the cube.
	square, _ = squareAndCube(5)
	fmt.Println(square)

	// Receive only the cube; ignore the square.
	_, cube = squareAndCube(3)
	fmt.Println(cube)

	// 10.
	sum, product, average := numberStats(2, 3, 4)
	fmt.Println("Sum:", sum, "Product:", product, "Average:", average)
}
