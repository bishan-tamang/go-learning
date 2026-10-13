package main

import "fmt"

func numberAnalyzer() {
	var number int

	fmt.Print("Enter a number:")
	_, err := fmt.Scan(&number)

	if err != nil {
		fmt.Println("Please enter a valid integer.")
		return
	}

	if number > 0 {
		fmt.Println("postive")
	} else if number < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("zero")
	}

	if number%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if number%3 == 0 && number%5 == 0 {
		fmt.Println("Divisible by both 3 and 5")
	} else {
		fmt.Println("Not Divisible by both 3 and 5")
	}
}

func main() {
	numberAnalyzer()
}
