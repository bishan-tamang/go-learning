package main

import "fmt"

func digitAnalyzer(number int) {
	digitCount := 0
	sumOfDigits := 0
	reversedNumber := 0

	// Use a temporary variable 'i' so we don't destroy the original number
	for i := number; i > 0; i /= 10 {
		lastDigit := i % 10 // 1. Get the rightmost digit

		digitCount++                                       // 2. Count it
		sumOfDigits += lastDigit                           // 3. Add to the total sum
		reversedNumber = (reversedNumber * 10) + lastDigit // 4. Build the reverse
	}

	fmt.Println("Digit Count:    ", digitCount)
	fmt.Println("Sum of Digits:  ", sumOfDigits)
	fmt.Println("Reversed Number:", reversedNumber)
}

func main() {
	var number int
	fmt.Print("Enter a positive integer: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error: Invalid input. Please enter numbers only.")
		return
	}

	if number <= 0 {
		fmt.Println("Error: Please enter a number greater than 0.")
		return
	}

	digitAnalyzer(number)
}
