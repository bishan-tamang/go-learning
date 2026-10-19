package main

import "fmt"

func multiplicationTableTool(number int) {
	var limit int

	fmt.Print("Enter a table limit: ")
	_, err := fmt.Scan(&limit)
	if err != nil {
		fmt.Println("Error: Enter a valid integer for limit.")
		return
	}

	// Reject a limit below 1.
	if limit < 1 {
		fmt.Println("Limit must be at least 1.")
		return
	}

	var order string

	fmt.Print("Enter an order (ascending/descending): ")
	_, err = fmt.Scan(&order)
	if err != nil {
		fmt.Println("Error: Could not read the order.")
		return
	}

	// Print the table in the chosen order.
	switch order {
	case "ascending":
		for i := 1; i <= limit; i++ {
			fmt.Printf("%d x %d = %d\n", number, i, number*i)
		}
	case "descending":
		for i := limit; i >= 1; i-- {
			fmt.Printf("%d x %d = %d\n", number, i, number*i)
		}
	default:
		fmt.Println("Error: Choose ascending or descending.")
	}
}

func main() {
	var number int

	fmt.Print("Enter a number for multiplication: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error: Please enter a valid integer for multiplication.")
		return
	}

	multiplicationTableTool(number)
}
