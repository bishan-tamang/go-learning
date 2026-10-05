package main

import "fmt"

// 1. Double a number
// Creating a function when number become double
func double(number int) int {
	return number * 2
}

// 2. Return a value without parameters
func favoriteLanguage() string {
	return "Go"
}

// 3. Return a Boolean
func isEven(num int) bool {
	return num%2 == 0
}

// 4. Use early returns
func numberLabel(number int) string {
	if number > 0 {
		return "positive"
	}

	if number < 0 {
		return "negative"
	}

	return "zero"
}

// 5. Return two integers
func calculate(a, b int) (int, int) {
	sum := a + b
	product := a * b

	return sum, product
}

// 6. Return different data types
func studentInfo() (string, int, bool) {
	name := "Bishan"
	age := 22
	isStudent := true

	return name, age, isStudent
}

// 7. Return a result and an error
func withdraw(balance, amount int) (int, error) {
	if amount <= 0 {
		return 0, fmt.Errorf("Amount must be greater then 0.")
	}

	if amount > balance {
		return 0, fmt.Errorf("amount exceeds balance.")
	}

	return balance - amount, nil
}

func main() {
	// 1. Store the result in a variable and printing it.
	doubleResult := double(5)
	fmt.Println(doubleResult)

	// 2. Calling the function, and store its return value, and print it.
	language := favoriteLanguage()
	fmt.Println(language)

	// 3. Return true when the number is even; otherwise, return false.
	fmt.Println(isEven(8))
	fmt.Println(isEven(7))

	// 4. Calling the use early return function
	fmt.Println(numberLabel(7))
	fmt.Println(numberLabel(-7))
	fmt.Println(numberLabel(0))

	// 5. Call the two intergers return value
	sum, product := calculate(7, 7)
	fmt.Println(sum, product)

	// Calling again but using _ to ignore the sum
	_, product = calculate(7, 6)
	fmt.Println(product)

	// 6. Calling studentInfo function with different return values
	studnetName, studentAge, studentStatus := studentInfo()
	fmt.Println(studnetName, studentAge, studentStatus)

	// 7.
	remaining, err := withdraw(1000, 300)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Remaining balance:", remaining)
}
