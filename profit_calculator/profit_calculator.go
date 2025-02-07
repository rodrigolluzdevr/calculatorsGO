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

// // Earnings Before Tax function
// func earningsBeforeTax(revenue float64, expenses float64) (ebt float64) {
// 	ebt = revenue - expenses
// 	return ebt
// }

// // Earnings After Tax function
// func earningsAfterTax(ebt float64, taxRate float64) (profit float64) {
// 	profit = ebt * (1 - taxRate/100)
// 	return profit
// }

// // Ratio function
// func ratio(ebt float64, profit float64) (ratio float64) {
// 	ratio = ebt / profit
// 	return ratio
// }
