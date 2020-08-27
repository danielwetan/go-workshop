package main

import "fmt"

// # Sales tax calculator

/*
	In this activity, we create a shopping cart application,
	where sales tax must be added to calculate the total.
	1. Create a calculator that calculates the sales tax for a single item
	1. The calculator must take the items cost and its sales tax rate
	3. Sum the sales tax and print the total amount of sales tax required for the following items
*/

func salesTax(cost, taxRate float64) float64 {
	return cost * taxRate
}

func main() {
	taxTotal := .0

	// Cake
	taxTotal += salesTax(.99, .075)

	// Milk
	taxTotal += salesTax(2.75, .015)

	// Butter
	taxTotal += salesTax(.87, .02)

	// Total
	fmt.Println("Sales tax total:", taxTotal)
}
