package main

import "fmt"

// Level 2: Combine operations

// Exercise 4: Bank Account Management
// Create a map containing account holders and their balances.
// Requirements:
// 1. Add a new account.
// 2. Update Bishan's balance.
// 3. Delete Sophia's account.
// 4. Print all account holders and balances.
func bankAccountManagement() {
	accounts := map[string]int{
		"Mamita": 1000,
		"Bishan": 999,
		"Sophia": 750,
	}

	// 1. Add a new account.
	accounts["Sonam"] = 650

	// 2. Update Bishan's balance.
	accounts["Bishan"] = 1100

	// 3. Delete Sophia's account.
	delete(accounts, "Sophia")

	// 4. Print all account holders and balances.
	fmt.Println("All account holders and balances:", accounts)
}

// Exercise 5: Shopping Cart
// Create a map containing product names and quantities.
// Requirements:
// 1. Add three products.
// 2. Update the quantity of one product.
// 3. Remove one product.
// 4. Add another product.
// 5. Print the final shopping cart.
func shoppingCart() {
	cart := make(map[string]int)

	// 1. Add three products.
	cart["Bag"] = 10
	cart["Pant"] = 25
	cart["Shirt"] = 15

	// 2. Update the quantity of one product.
	cart["Shirt"] = 20

	// 3. Remove one product.
	delete(cart, "Bag")

	// 4. Add another product.
	cart["Shoes"] = 15

	// 5. Print the final shopping cart.
	fmt.Println("The final shopping cart:", cart)
}

// Exercise 6: Employee Management
// Create a map containing employee names and salaries.
// Requirements:
// 1. Add four employees.
// 2. Update two employees' salaries.
// 3. Delete one employee.
// 4. Add a new employee.
// 5. Print the final employee records.
func employeeManagement() {
	employees := make(map[string]int)

	// 1. Add four employees.
	employees["Ram"] = 100
	employees["Hari"] = 80
	employees["Sita"] = 60
	employees["Gita"] = 40

	// 2. Update two employees' salaries.
	employees["Sita"] = 75
	employees["Gita"] = 70

	// 3. Delete one employee.
	delete(employees, "Gita")

	// 4. Add a new employee.
	employees["Raima"] = 88

	// 5. Print the final employee records.
	fmt.Println("The final employee records:", employees)
}

func main() {
	bankAccountManagement()
	shoppingCart()
	employeeManagement()
}
