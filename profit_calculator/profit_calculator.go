package main

import (
	"fmt"
)

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	outputText("Revenue: ")
	inputText(&revenue)

	outputText("Expenses: ")
	inputText(&expenses)

	outputText("Tax Rate: ")
	inputText(&taxRate)

	// variables def
	ebt, profit, ratio := calculateTax(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax - EBT: %0.2f\n", ebt)

	fmt.Printf("Earnings After Tax - profit: %0.2f\n", profit)

	fmt.Printf("Ratio: %0.2f\n", ratio)

}

// calculate tax in earnings before and after tax, ratio.
func calculateTax(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

// output function
func outputText(text string) {
	fmt.Print(text)
}

// input function
func inputText(variable *float64) {
	fmt.Scan(variable)
}
