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
}
