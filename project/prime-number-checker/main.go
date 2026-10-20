package main

import "fmt"

func primeNumberChecker(number int) {
	// Prime numbers must be greater than 1.
	if number < 2 {
		fmt.Println("The number is not prime.")
		return
	}

	isPrime := true

	// Look for a divisor other than 1 and the number itself.
	for divisor := 2; divisor < number; divisor++ {
		if number%divisor == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println("The number is prime.")
	} else {
		fmt.Println("The number is not prime.")
	}
}

func main() {
	var number int

	fmt.Print("Enter an integer: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Enter a valid integer number.")
		return
	}

	primeNumberChecker(number)
}
