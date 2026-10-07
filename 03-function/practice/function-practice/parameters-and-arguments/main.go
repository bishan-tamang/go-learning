package main

import "fmt"

// 1. Personal greeting: Write greet(name string) to print "Hello, Bishan!" when called with "Bishan". Call it with three different names.
func greet(name string) {
	fmt.Printf("Hello, %s\n", name)
}

// 2. Display an age: Write showAge(age int) to print "You are 22 years old." when given 22.
func showAge(age int) {
	fmt.Println("You are", age, "Years old.")
}

// 3. Square a number: Write printSquare(number int) to print the square of the argument. Test with 5, -3, and 0.
func printSquare(number int) {
	result := number * number
	fmt.Println(result)
}

// 4. Check even or odd: Write checkEvenOdd(number int) to print whether the argument is even or odd.
func checkEvenOdd(number int) {
	if number%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

// 5. Use different parameter types: Write introduce(name string, age int) to print a sentence containing both values. Call it twice with different arguments.
func introduce(name string, age int) {
	fmt.Println("My name is", name)
	fmt.Println("I am", age, "years old")
}

// 6. Add two numbers: Write printSum(a int, b int) to print their sum. Call it once with literal arguments, such as printSum(10, 20), and once with variables.
func printSum(a int, b int) {
	fmt.Println(a + b)
}

// 7. Repeat a message: Write repeatMessage(message string, times int) to print the message times times using a loop. If times <= 0, print nothing.
func repeatMessage(message string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(message)
	}
}

// 8. Multiplication table: Write printTable(number int, limit int) to print the multiplication table of number from 1 through limit. Test with (7, 10) and (3, 5).
func printTable(number int, limit int) {
	for i := 1; i <= limit; i++ {
		result := number * i
		fmt.Println(result)
	}
}

// 9. Predict the output: Identify the parameters and arguments, then predict the output before running:
// package main

// import "fmt"

// func showDetails(name string, age int) {
//     fmt.Println(name, age)
// }

// func main() {
//     person := "Bishan"
//     years := 22

//	    showDetails(person, years)
//	    showDetails("Sita", 25)
//	}

// Outpu:
// Bishan 22
// Sita 25
// Parameters are: (name string, age int)
// Arguments are: (person, years), ("site", 25)

// 10. Fix incorrect calls: Given this function:
func introduces(name string, age int) {
	fmt.Println(name, age)
}

// Explain what is wrong with each call, then correct it:
// introduce(22, "Bishan")
// introduce("Bishan")
// introduce("Bishan", "22")
// introduce("Bishan", 22, "Nepal")

// There is argument is not match to parameters, in first there is int and stirng but i should be string and int, and in second call there is only one argument but parameters need two arguments, in third call there is both in string, and in fourth call there is three arguments but parameters need only two.

func main() {
	// 1.
	greet("Bishan")
	greet("Mamita")
	greet("Ismita")

	// 2.
	showAge(22)

	// 3.
	printSquare(5)
	printSquare(-3)
	printSquare(0)

	// 4.
	checkEvenOdd(7)
	checkEvenOdd(8)

	// 5.
	introduce("Bishan", 22)
	introduce("Mamita", 21)

	// 6.
	// Call with literal arguments
	printSum(10, 20)

	// Call with variables as arguments
	num1 := 50
	num2 := 70
	printSum(num1, num2)

	// 7.
	repeatMessage("I love Go!", 3)
	repeatMessage("Hello!", 0)
	repeatMessage("Hello!", -2)

	// 8.
	printTable(7, 10)
	printTable(3, 5)

	// 10.
	// // There is the correct version:
	introduces("Bishan", 22)
}
