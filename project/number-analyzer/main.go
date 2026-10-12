package main

import "fmt"

func numberAnalyzer() {
	var number int

	fmt.Print("Enter a number:")
	fmt.Scan(&number)

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

	if number%5 == 0 {
		fmt.Println("Divisible by 5")
	} else {
		fmt.Println("Not Divisible by 5")
	}
}

func main() {
	numberAnalyzer()
}
