package main

import "fmt"

// 11. Create three functions:
// 	. morning()
// 	. afternoon()
// 	. evening()

// Call all three from main().

func morning() {
	fmt.Println("Good Morning")
}

func afternoon() {
	fmt.Println("Good Afternoon")
}

func evening() {
	fmt.Println("Good Evening")
}

// 12. Create four functions:
// 	. name()
// 	. age()
// 	. country()
// 	. language()

// Call them from main().

func name() {
	fmt.Println("My name is Bishan Tamang")
}

func age() {
	fmt.Println("I am 22 years old.")
}

func country() {
	fmt.Println("I live in Nepal")
}

func language() {
	fmt.Println("I speak Nepali language")
}

// 13. Create a function called start() that prints:
// Starting program...

// Create another function called finish() that prints:
// Program finished.

// Call both from main().

func start() {
	fmt.Println("Starting program...")
}

func finish() {
	fmt.Println("Program finished.")
}

// 14. Create a function called greetMother() and another called greetFather(). Each should print a different greeting.
func greetMother() {
	fmt.Println("Hello, Mother!")
}

func greetFather() {
	fmt.Println("Hello, Father!")
}

// 15. Create three functions:
// 	. learn()
// 	. practice()
// 	. build()

// Make them print:
// 	. I am learning
// 	. I am practicing
// 	. I am building

// Call them in that exact order.

func learn() {
	fmt.Println("I am learning")
}

func practice() {
	fmt.Println("I am practicing")
}

func build() {
	fmt.Println("I am building")
}

// Now calling all function in main function
func main() {
	morning()
	afternoon()
	evening()
	name()
	age()
	country()
	language()
	start()
	finish()
	greetMother()
	greetFather()
	learn()
	practice()
	build()
}
