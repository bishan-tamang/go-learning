package main

import "fmt"

// Exercise 6: User Login System
func userLoginSystem() {
	users := map[string]string{
		"bishan": "12345",
		"ram":    "ram123",
		"sita":   "sita456",
	}

	var username string

	fmt.Print("Enter username: ")
	if _, err := fmt.Scan(&username); err != nil {
		fmt.Println("Invalid input! Please try again.")
		return
	}

	// Check whether the username exists.
	storedPassword, ok := users[username]

	if !ok {
		fmt.Println("Username not found.")
		return
	}

	var password string

	fmt.Print("Enter password: ")
	if _, err := fmt.Scan(&password); err != nil {
		fmt.Println("Invalid input! Please try again.")
		return
	}

	// Check whether the password matches.
	if password != storedPassword {
		fmt.Println("Incorrect password.")
		return
	}

	fmt.Println("Login successful!")
}

// Exercise 7: Employee Lookup
func employeeLookup() {
	employees := map[int]string{
		101: "Bishan",
		102: "Ram",
		103: "Sita",
	}

	// Search employee ID 102.
	employee102, ok := employees[102]

	if ok {
		fmt.Printf("Employee ID 102: %s\n", employee102)
	} else {
		fmt.Println("Employee ID 102 not found.")
	}

	// Search employee ID 105.
	employee105, ok := employees[105]

	if ok {
		fmt.Printf("Employee ID 105: %s\n", employee105)
	} else {
		fmt.Println("Employee ID 105 not found.")
	}
}

// Exercise 8: Shopping Cart
func shoppingCart() {
	cart := map[string]int{
		"Apple":  5,
		"Banana": 0,
		"Mango":  3,
	}

	// Check Banana.
	bananaQuantity, ok := cart["Banana"]

	if !ok {
		fmt.Println("Banana is not in the cart.")
	} else {
		fmt.Printf("Banana quantity: %d\n", bananaQuantity)
	}

	// Check Orange.
	orangeQuantity, ok := cart["Orange"]

	if !ok {
		fmt.Println("Orange is not in the cart.")
	} else {
		fmt.Printf("Orange quantity: %d\n", orangeQuantity)
	}
}

// Exercise 9: Update Marks Safely
func updateMarksSafely() {
	marks := map[string]int{
		"Bishan": 80,
		"Ram":    75,
	}

	// Update Bishan's marks.
	if _, ok := marks["Bishan"]; !ok {
		fmt.Println("Cannot update: Bishan not found.")
	} else {
		marks["Bishan"] = 95
		fmt.Println("Bishan's updated mark:", marks["Bishan"])
	}

	// Update Sita's marks.
	if _, ok := marks["Sita"]; !ok {
		fmt.Println("Cannot update: Sita not found.")
	} else {
		marks["Sita"] = 85
		fmt.Println("Sita's updated mark:", marks["Sita"])
	}

	fmt.Println("Final marks:", marks)
}

// Exercise 10: Count Product Availability
func countProductAvailability() {
	products := map[string]int{
		"Laptop": 10,
		"Mouse":  0,
		"Phone":  5,
		"Tablet": 0,
	}

	var productName string

	fmt.Print("Enter a product name: ")
	if _, err := fmt.Scan(&productName); err != nil {
		fmt.Println("Invalid input! Please try again.")
		return
	}

	// Check whether the product exists.
	quantity, ok := products[productName]

	if !ok {
		fmt.Println("Product not found.")
		return
	}

	// Check product availability.
	if quantity > 0 {
		fmt.Printf("%s is in stock. Quantity: %d\n", productName, quantity)
	} else {
		fmt.Printf("%s is out of stock.\n", productName)
	}
}

func main() {
	userLoginSystem()
	employeeLookup()
	shoppingCart()
	updateMarksSafely()
	countProductAvailability()
}
