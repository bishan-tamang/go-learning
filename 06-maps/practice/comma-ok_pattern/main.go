package main

import "fmt"

// Exercise 1: Check Student Name
func checkStudentName() {
	marks := map[string]int{
		"Bishan": 85,
		"Ram":    90,
		"Sita":   0,
	}

	// Check whether Bishan exists.
	mark, ok := marks["Bishan"]

	if ok {
		fmt.Println("Mark:", mark)
	} else {
		fmt.Println("Student not found")
	}
}

// Exercise 2: Find a Missing Student
func missingStudent() {
	marks := map[string]int{
		"Bishan": 88,
		"Ram":    90,
		"Sita":   0,
	}

	// Check whether Hari exists.
	_, ok := marks["Hari"]

	if ok {
		fmt.Println("Student found")
	} else {
		fmt.Println("Student not found")
	}
}

// Exercise 3: Understand Zero Values
func understandingZeroValues() {
	stock := map[string]int{
		"Laptop": 10,
		"Mouse":  0,
	}

	// Check Mouse.
	mouseQuantity, ok := stock["Mouse"]

	fmt.Println("Mouse quantity:", mouseQuantity)
	fmt.Println("Mouse exists:", ok)

	if ok {
		fmt.Println("Product found!")
	} else {
		fmt.Println("Product not found!")
	}

	fmt.Println()

	// Check Keyboard.
	keyboardQuantity, ok := stock["Keyboard"]

	fmt.Println("Keyboard quantity:", keyboardQuantity)
	fmt.Println("Keyboard exists:", ok)

	if ok {
		fmt.Println("Product found!")
	} else {
		fmt.Println("Product not found!")
	}
}

// Exercise 4: Check User Account
func checkUserAccount() {
	users := map[string]bool{
		"Ram":  true,
		"Sita": false,
	}

	// Check Ram.
	ramActive, ok := users["Ram"]

	if ok {
		fmt.Println("Ram account exists:", ramActive)
	} else {
		fmt.Println("Ram account does not exist.")
	}

	// Check Sita.
	sitaActive, ok := users["Sita"]

	if ok {
		fmt.Println("Sita account exists:", sitaActive)
	} else {
		fmt.Println("Sita account does not exist.")
	}

	// Check Hari.
	hariActive, ok := users["Hari"]

	if ok {
		fmt.Println("Hari account exists:", hariActive)
	} else {
		fmt.Println("Hari account does not exist.")
	}

	// Explanation:
	// The stored boolean value tells us whether the account is active.
	// The ok value tells us whether the key exists in the map.
}

// Exercise 5: Find a Product Price
func findProductPrice() {
	prices := map[string]float64{
		"Laptop": 850.50,
		"Mouse":  25.50,
		"Phone":  500.00,
	}

	var productName string

	fmt.Print("Enter a product name: ")
	if _, err := fmt.Scan(&productName); err != nil {
		fmt.Println("Invalid input! Please try again.")
		return
	}

	price, ok := prices[productName]

	if !ok {
		fmt.Println("Product not found")
		return
	}

	fmt.Printf("%s : $%.2f\n", productName, price)
}

func main() {
	checkStudentName()
	missingStudent()
	understandingZeroValues()
	checkUserAccount()
	findProductPrice()
}
