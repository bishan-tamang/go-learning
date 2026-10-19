package main

import "fmt"

func analyzeNumber(number int) (string, string, string) {
	var sign, parity, divisibility string

	// Check whether the number is positive, negative, or zero.
	if number > 0 {
		sign = "positive"
	} else if number < 0 {
		sign = "negative"
	} else {
		sign = "zero"
	}

	// Check whether the number is even or odd.
	if number%2 == 0 {
		parity = "even"
	} else {
		parity = "odd"
	}

	// Check divisibility by both 3 and 5.
	if number%3 == 0 && number%5 == 0 {
		divisibility = "yes"
	} else {
		divisibility = "no"
	}

	return sign, parity, divisibility
}

func main() {
	var number int

	fmt.Print("Enter an integer: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Please enter a valid integer.")
		return
	}

	sign, parity, divisibility := analyzeNumber(number)

	fmt.Println("Sign:", sign)
	fmt.Println("Parity:", parity)
	fmt.Println("Divisible by both 3 and 5:", divisibility)
}
