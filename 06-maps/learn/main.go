package main

import "fmt"

// 1. What is maps?
// A map is a built-in data structure that stores data in key-value pairs, allowing us to access, add, update, and delete values using their keys.
// Think of a map like a dictionary: you use a key to find its corresponding values.

// For example, imagine storing student information:
// Key			Value
// name			Bishan
// age 			22
// city			Lalitpur

// Here:
// . Key: used to identify the data
// . Vlaue: The actual data associated with that key.

// Here's the code:
func sampleMaps() {
	studentInformation := map[string]string{
		"Name": "Bishan",
		"Age":  "22",
		"City": "Lalitpur",
	}

	fmt.Println(studentInformation)
}

// Understandt this declaration:
// map(string)int

// . map - creates a map type.
// . string - the key type.
// . string - the value type.

// This map stores string keys and string values.

// Other common map types include:
func otherCommonMapTypes() {
	// string and float64
	price := map[string]float64{
		"laptop":   850.50,
		"Keyboard": 45.00,
	}

	// int and string
	userIds := map[int]string{
		101: "Bishan",
		102: "Mamita",
	}

	// string and boolean
	activeUser := map[string]bool{
		"Ram":  true,
		"Sita": false,
	}

	fmt.Println(price)
	fmt.Println(userIds)
	fmt.Println(activeUser)
}

// Real-world usage: map[string]User for looking up users by username, or map[string]int for counting occurrences.

// 2. Creating maps - the three patterns you need
// Pattern A: Create and initialize a map
// Use this when you already know the initial values.
func createAndInitializeAMap() {
	ages := map[string]int{
		"Ram":  22,
		"Hari": 23,
	}

	fmt.Println(ages)
}

// Pattern B: Create an empty map using make
// Use this when you'll add data later.
func usingMake() {
	ages := make(map[string]int)
	fmt.Println("Before adding keys and values:", ages)

	ages["Bishan"] = 22
	ages["Mamita"] = 21

	fmt.Println("After adding keys and values:", ages)
}

// // An empty map created with make is ready for assignments.

// Pattern C: Declare a nil map
func nilMaps() {
	var ages map[string]int

	fmt.Println(nil == ages) // Output: true
	fmt.Println(len(ages))   // Output: 0

	fmt.Println(ages)
}

// You can read from a nil map, but you cannot assign new entries to it.

// Let's try to assign nil map
func assignNillMap() {
	var ages map[string]int

	// ages["Ram"] = 22

	fmt.Println(ages)

	// Output: // PANIC: assignment to entry in nil map
}

// Remember this distinction:
// . nil map: can be read but cannot accept assignments.
// . Initialized empty map: can be read and modified.

// To fix it, initialize the map first using map literal or make.
// For most practical situations, use a map literal or make.
func main() {
	sampleMaps()
	otherCommonMapTypes()
	createAndInitializeAMap()
	usingMake()
	nilMaps()
	assignNillMap()
}
