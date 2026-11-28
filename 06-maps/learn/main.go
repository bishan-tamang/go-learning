package main

import (
	"fmt"
	"sort"
)

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

// 3. Add, update, read and delete - the four essential operations.
// There are the operations you'll used most ofthen.

// A. Add a new key-value pari
func addNewKeyValue() {
	// Using make
	ages := make(map[string]int)

	ages["Ram"] = 21
	ages["Hari"] = 22

	fmt.Println(ages)

	// If a key is not in the map, Go automatically adds it with the given value.

	// Using map literal
	userIds := map[int]string{
		101: "Bishan",
		102: "Mamita",
	}

	userIds[103] = "Sophia"

	fmt.Println(userIds)
}

// B. Update an existing value
func updateAnExistingValue() {
	prices := map[string]float64{
		"Laptop": 499.99,
		"PS5":    999.99,
	}

	fmt.Println("Before update an exixting value: ", prices)

	prices["PS5"] = 799.00
	fmt.Println("After update an existing vlaue: ", prices)

	prices["Mouse"] = 49.99 // Creates a new entry
	fmt.Println(prices)
}

// Remember:
// . If PS5 already exists, its value is replaced. If it doesn't exist, the assignment creates a new entry (one complete key-value pair).

// Remember:
// . Entry means One complete key-value pair in a map.

// C. Read/Access a value from map
func accessAValueFromMap() {
	// Studnet mark
	marks := map[string]int{
		"Bishan": 99,
		"Mamita": 98,
	}

	fmt.Println("The mark of Bishan is", marks["Bishan"]) // Output: 99

	// What if the key doen't exit?
	fmt.Println(marks["Sophia"]) // Output: 0
}

// Because the value type is int, So Go returns it zero vlaue, 0, when the key doen't exist. It doen't throw an error.
// This creates an important problem: how do you distinguish a missing key from a key whose value is actually 0?
// We'll solve that next.

// D. Delete an entry
func deleteAnEntry() {
	ages := map[string]int{
		"Mamita": 21,
		"Bishan": 22,
		"Hari":   23,
	}

	fmt.Println("Before delete an entry:", ages)

	delete(ages, "Hari")
	fmt.Println("Afte delete an entry:", ages)
	// If key exists, it's removed.

	delete(ages, "Sita")
	fmt.Println("Delete an non-existing entry:", ages)
	// If doesn't exist, delete does nothing.
}

// There four operations from the foundation of almost evey map-based program.

// 5. The comma-ok pattern - one of the most important Go pattern
// Suppose you have this map.
func commaOk() {
	stock := map[string]int{
		"Laptop": 10,
		"Mouse":  0,
	}

	// If you access a missing key, the result is 0.
	fmt.Println(stock["Keyboard"]) // Output: 0

	// But "Mouse" also has a stock value of 0.
	fmt.Println(stock["Mouse"]) // Output: 0
}

// Use the comma-ok pattern to check whether the key exits.
func commaOkPattern() {
	stock := map[string]int{
		"Laptop": 10,
		"Mouse":  0,
	}

	quantity, ok := stock["Mouse"]

	if ok {
		fmt.Println("Stock:", quantity)
	} else {
		fmt.Println("Product not found!")
	}

	// Output: Stock: 0
}

// Now try a missing product:
func missingProduct() {
	stock := map[string]int{
		"Laptop": 10,
		"Mouse":  0,
	}

	quantity, ok := stock["Keyboard"]

	if !ok {
		fmt.Println("Product not found")
	} else {
		fmt.Println("Stock:", quantity)
	}

	// Output: Product not found
}

// The key idea is:
// value, ok := myMap[key]
// . value contains the value stored in the map
// . ok tells you whether the key exist (true) or doen't exist (false).
// . If the key doesn't exist, value gets the zero value of its data type.
// . Even if the stored value is 0, false, or "", ok will still be true if the key exist.

// Remember: value tells you what is stored, and ok tells you whether the key exist.

// A real-world example: user lookup
func realWorldExample() {
	users := map[int]string{
		101: "Ram",
		102: "Sita",
	}

	id := 103

	name, ok := users[id]

	if !ok {
		fmt.Println("User not found")
		return
	}

	fmt.Println(name)
}

// 6. Looping through maps with range
// You can also range to process every entry in a map.
func loopingWithRange() {
	marks := map[string]int{
		"Ram":  85,
		"Sita": 92,
		"Hari": 78,
	}

	// Loop through range
	for studentName, mark := range marks {
		fmt.Println(studentName, mark)
	}
}

// This print each student name and mark, but the order is not guaranted.
// For example, you might see Ram first in one run and Hari in another. Don't rely on map iteration order.

// loop through only the keys
func loopThroughOnlyKeys() {
	marks := map[string]int{
		"Ram":  85,
		"Sita": 92,
		"Hari": 78,
	}

	// Loop through only the keys
	for name := range marks {
		fmt.Println("Only key:", name)
	}
}

// Loop through only the values
func loopThroughOnlyValues() {
	marks := map[string]int{
		"Ram":  85,
		"Sita": 92,
		"Hari": 78,
	}

	// Loop through only the values
	for _, mark := range marks {
		fmt.Println("Only value:", mark)
	}

	// The underscore _ ignores the key.
}

// When you need to sorted output
// If you need a consistent order, such an in reports or tests, collect the keys and sort them.
// In Go, you can not directly sort a map. Go maps do not maintain a guaranteed order.

// If you want sorted output from a a map, you should:
// 1. Collect the map keys into a slice.
// 2. Sort the slice
// 3. Loop through the sorted slice
// 4. Access the map values using those keys.

func sortedOutput() {
	marks := map[string]int{
		"Ram":  85,
		"Sita": 92,
		"Hari": 78,
	}

	names := make([]string, 0, len(marks))

	for name := range marks {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Println(name, marks[name])
	}
}

// Now the output is correctly ordered alphabetically by name.

// Remember: Maps are excellent for lookup. Slice are useful when you need ordering. In Go, you'll often use them together.

func main() {
	sampleMaps()
	otherCommonMapTypes()
	createAndInitializeAMap()
	usingMake()
	nilMaps()
	assignNillMap()
	addNewKeyValue()
	updateAnExistingValue()
	accessAValueFromMap()
	deleteAnEntry()
	commaOk()
	commaOkPattern()
	missingProduct()
	realWorldExample()
	loopingWithRange()
	loopThroughOnlyKeys()
	loopThroughOnlyValues()
	sortedOutput()
}
