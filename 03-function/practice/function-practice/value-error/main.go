package main

import (
	"errors"
	"fmt"
)

// 1. Divide two numbers
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divided by zero")
	}

	return a / b, nil
}

// 2. Validate an age
func validateAge(age int) (int, error) {
	if age < 0 || age > 120 {
		return 0, errors.New("age must be between 0 and 120")
	}
	return age, nil
}

// 3. Calculate a rectangle’s area
func rectangleArea(lenght, width float64) (float64, error) {
	if lenght <= 0 || width <= 0 {
		return 0, errors.New("Both dimensions must be greater than zero")
	}

	return lenght * width, nil
}

// 4. Validate a name
func validateName(name string) (string, error) {
	if name == "" {
		return name, errors.New("name must not be empty")
	}

	return name, nil
}

// 5. Withdraw money
func WithdrawMoney(balance, amount int) (int, error) {
	if balance < 0 {
		return 0, errors.New("balance must not be negative")
	}

	if amount <= 0 {
		return 0, errors.New("withdrawal amount must be greater than zero")
	}

	if amount > balance {
		return 0, errors.New("insufficient balance")
	}

	return balance - amount, nil
}

// 6. Calculate a discounted price
func discountPrice(price, discountPercent float64) (float64, error) {
	if price < 0 || discountPercent < 0 || discountPercent > 100 {
		return 0, errors.New("The price must be nonnegative, and discount must be between 0 and 100")
	}

	finalPrice := price - (price * discountPercent / 100)
	return finalPrice, nil
}

// 7. Convert a score to a grade
func grade(score int) (string, error) {
	if score < 0 || score > 100 {
		return "", errors.New("Invalid Grade")
	}

	var grade string

	if score >= 90 {
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

	return grade, nil
}

// 8. Calculate a factorial
func factorial(n int) (int, error) {
	if n < 0 || n > 10 {
		return 0, errors.New("Accept only values from 0 through 10.")
	}

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial = factorial * i
	}

	return factorial, nil
}

// 9. Build a calculator function
func calculator(a, b float64, operation string) (float64, error) {
	switch operation {
	case "add":
		return a + b, nil
	case "subtract":
		return a - b, nil
	case "multiply":
		return a * b, nil
	case "divide":
		if b == 0 {
			return 0, errors.New("Cannot divide by zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("Unknown operation")
	}
}

// 10. Call a function that can fail
func average(total float64, count int) (float64, error) {
	if count <= 0 {
		return 0, errors.New("count must be greater than zero")
	}

	result, err := divide(total, float64(count))
	if err != nil {
		return 0, err
	}

	return result, nil
}

func main() {
	// 1.
	value, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", value)
	}

	value, err = divide(7, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", value)
	}

	value, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", value)
	}

	// 2.
	age, err := validateAge(22)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age:", age)
	}

	age, err = validateAge(0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age:", age)
	}

	age, err = validateAge(120)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age:", age)
	}

	age, err = validateAge(-1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age:", age)
	}

	age, err = validateAge(121)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age:", age)
	}

	// 3.
	area, err := rectangleArea(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Area of Rectangle:", area)
	}

	area, err = rectangleArea(0, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Area of Rectangel:", area)
	}

	area, err = rectangleArea(10, -12)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Area of Rectangel:", area)
	}

	area, err = rectangleArea(-3, -4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Area of Rectangle:", area)
	}

	// 4.
	name, err := validateName("Bishan")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Name:", name)
	}

	name, err = validateName("")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Name:", name)
	}

	// 5.
	remaining, err := WithdrawMoney(1000, 300)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Remaining balance:", remaining)
	}

	remaining, err = WithdrawMoney(1000, 1500)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Remaining balance:", remaining)
	}

	remaining, err = WithdrawMoney(1000, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Remaining balance:", remaining)
	}

	remaining, err = WithdrawMoney(-100, 50)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Remaining balance:", remaining)
	}

	// 6.
	finalPrice, err := discountPrice(1000, 20)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Final Price:", finalPrice)
	}

	finalPrice, err = discountPrice(1000, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Final Price:", finalPrice)
	}

	finalPrice, err = discountPrice(-100, 20)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Final Price:", finalPrice)
	}

	finalPrice, err = discountPrice(1000, -5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Final Price:", finalPrice)
	}

	finalPrice, err = discountPrice(1000, 101)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Final Price:", finalPrice)
	}

	// 7.
	grades, err := grade(59)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Grade:", grades)
	}

	grades, err = grade(60)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Grade:", grades)
	}

	grades, err = grade(90)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Grade:", grades)
	}

	grades, err = grade(-1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Grade:", grades)
	}

	grades, err = grade(101)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Grade:", grades)
	}

	// 8.
	factorials, err := factorial(0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Factorial:", factorials)
	}

	factorials, err = factorial(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Factorial:", factorials)
	}

	factorials, err = factorial(-1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Factorial:", factorials)
	}

	factorials, err = factorial(11)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Factorial:", factorials)
	}

	// 9.
	result, err := calculator(6, 2, "add")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = calculator(6, 2, "subtract")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = calculator(6, 2, "multiply")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = calculator(6, 2, "divide")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = calculator(6, 0, "divide")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = calculator(6, 2, "unknown")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	results, err := average(100, 4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Average:", results)
	}

	results, err = average(5, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Average:", results)
	}

	results, err = average(100, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Average:", results)
	}

	results, err = average(100, -2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Average:", results)
	}
}
