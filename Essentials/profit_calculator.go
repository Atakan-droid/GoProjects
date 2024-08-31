package main

import "fmt"

func profit_calculator() {
	var revenue int
	var expenses int
	var taxRate float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := float64(ebt) * (1 - taxRate/100)
	ratio := ebt / int(profit)

	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}
