package main

// 1. What is Array?
// An array is a collection of a fixed number of values of the same type.

// Imagine you want to store three exam marks.
// Without an array, you might write.

mark1 := 70
mark2 := 85
mark3 := 90

// With an array, you can store them together:

var marks [3]int = [3]int{70,85,90}

// Here is what each part means:
// Code				Meaning
// var				Declears a variable
// marks 			Array ariable name
// [3]				Exactly three elements
// int 				Every elements stores an integer
// {70, 85, 90}		Initial values

// You can also use the shorter declaration:
marks := [3]int{70 , 85, 90}

// Both create an array containing three integers.

// How an array stores values
// Array: marks := [3]int{70, 85, 90}
// 			70			85			90
// 		  index 0	  index 1	  index 2
// There are 3 elements, and their indexes start at 0.

// Remember these two rules:
//	. An aaray has a fixed length.
//	. All elements in an array must have the same data types

// For example, [3] int can hold three integers, but it can't hold a string as one of its elements.
// Like this:
marks := [3]int{"Pass", 85, 90}

// Practice 1
// Try these question before moving to declaring arrays in Go.