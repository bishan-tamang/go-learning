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

func main() {
	sampleMaps()
	otherCommonMapTypes()
}
