package main

import "fmt"

// Exercise 7: Inventory Update System

// You have:
// inventory := map[string]int{
//     "Laptop":  10,
//     "Mouse":   20,
//     "Monitor": 5,
// }

// Requirements:
// - Add 5 laptops to the existing quantity.
// - Remove 3 mouse from the existing quantity.
// - Add 10 keyboards.
// - Delete monitors.
// - Print the final inventory.
func inventoryUpdateSystem() {
	inventory := map[string]int{
		"Laptop":  10,
		"Mouse":   20,
		"Monitor": 5,
	}

	// - Add 5 laptops to the existing quantity.
	inventory["Laptop"] += 5

	// - Remove 3 mouse from the existing quantity.
	inventory["Mouse"] -= 3

	// - Add 10 keyboards.
	inventory["Keyboard"] = 10

	// - Delete monitors.
	delete(inventory, "Monitor")

	// -  Print the final inventory.
	fmt.Println("The final inventory:", inventory)

}

// Exercise 8: Student Record Modification

// You have:
// students := map[string]int{
//     "Bishan": 85,
//     "Ram":    75,
//     "Hari":   90,
// }

// Requirements:
// - Change Bishan's marks to 95.
// - Add Sita with 88 marks.
// - Delete Ram.
// - Update Hari's marks to 92.
// - Print all students.
func studentRecordModification() {
	students := map[string]int{
		"Bishan": 85,
		"Ram":    75,
		"Hari":   90,
	}

	// - Change Bishan's marks to 95.
	students["Bishan"] = 95

	// - Add Sita with 88 marks.
	students["Sita"] = 88

	// - Delete Ram.
	delete(students, "Ram")

	// - Update Hari's marks to 92.
	students["Hari"] = 92

	// - Print all students.
	fmt.Printf("All students: %v\n", students)
}

// Exercise 9: Product Price Manager
// Create a map of products and prices.
// Requirements:
// - Add five products.
// - Change the price of two products.
// - Remove one product.
// - Add a new product.
// - Print all products and prices.
func productPriceManager() {
	products := make(map[string]float64)

	// - Add five products.
	products["USB"] = 45.99
	products["Ear Phone"] = 20.25
	products["Mouse"] = 45.84
	products["Mobile"] = 99.99
	products["Charger"] = 50.50

	// - Change the price of two products.
	products["USB"] = 5.90
	products["Ear Phone"] = 2.99

	// - Remove one product.
	delete(products, "Charger")

	// - Add a new product.
	products["Laptop"] = 599.00

	// - Print all products and prices.
	fmt.Println("All products and prices:", products)
}

func main() {
	inventoryUpdateSystem()
	studentRecordModification()
	productPriceManager()
}
