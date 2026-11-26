package main

import "fmt"

// Exercise 1: Personal Information
// Create a map containing:
// - Name
// - Age
// - City
// - Profession
// Requirements:
// 1. Print all information.
// 2. Update your city.
// 3. Update your profession.
// 4. Print the updated information.
func personalInfromation() {
	personalInfo := map[string]string{
		"Name":       "Bishan Tamang",
		"Age":        "22",
		"City":       "Kathmandu",
		"Profession": "Junior Backend Developer",
	}

	// 1. Print all information.
	fmt.Println("All infromation:", personalInfo)

	// 2. Update your city.
	personalInfo["City"] = "Lalitpur"

	// 3. Update your profession.
	personalInfo["Profession"] = "Senior Backend Enginner"

	// 4. Print the updated information.
	fmt.Println("Updated infromation:", personalInfo)
}

// Exercise 2: Student Management
// Create a map containing three students and their marks.
// Requirements:
// 1. Add Sita with 88 marks.
// 2. Update Ram's marks to 80.
// 3. Delete Hari.
// 4. Print all student records.
func studentManagement() {
	student := map[string]int{
		"Bishan": 85,
		"Ram":    75,
		"Hari":   90,
	}

	// 1. Add Sita with 88 marks.
	student["Sita"] = 88

	// 2. Update Ram's marks to 80.
	student["Ram"] = 80

	// 3. Delete Hari.
	delete(student, "Hari")

	// 4. Print all student records.
	fmt.Println("All student records:", student)
}

// Exercise 3: Product Inventory
// Create a map containing product names and quantities.
// Requirements:
// 1. Add Monitor with quantity 5.
// 2. Update Mouse quantity to 25.
// 3. Delete Keyboard.
// 4. Print the final inventory.
func productInventory() {
	inventory := map[string]int{
		"Laptop":  10,
		"Mouse":   20,
		"Keyword": 15,
	}

	// 1. Add Monitor with quantity 5.
	inventory["Monitor"] = 5

	// 2. Update Mouse quantity to 25.
	inventory["Mouse"] = 25

	// 3. Delete Keyboard.
	delete(inventory, "Keyword")

	// 4. Print the final inventory.
	fmt.Println("The final inventory:", inventory)
}

func main() {
	personalInfromation()
	studentManagement()
	productInventory()
}
