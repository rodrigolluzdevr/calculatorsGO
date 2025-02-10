package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	ultimateValueEbt    = "ultimateValueEbt.txt"
	ultimateValueProfit = "ultimateValueProfit.txt"
	ultimateValueRatio  = "ultimateValueRatio.txt"
)

// EBT SAVE
func writeUltimateValueEbt(ebt float64) error {
	ultimateValueEbtText := fmt.Sprint(ebt)
	err := os.WriteFile(ultimateValueEbt, []byte(ultimateValueEbtText), 0644)
	if err != nil {
		return fmt.Errorf("failed to save ebt: %w", err)
	}
	return nil
}

func getUltimateValueEbt() float64 {
	data, _ := os.ReadFile(ultimateValueEbt)
	ultimateValueEbtText := string(data)
	ebt, _ := strconv.ParseFloat(ultimateValueEbtText, 64)
	return ebt
}

// PROFIT SAVE
func writeUltimateValueProfit(profit float64) error {
	ultimateValueProfitText := fmt.Sprint(profit)
	err := os.WriteFile(ultimateValueProfit, []byte(ultimateValueProfitText), 0644)
	if err != nil {
		return fmt.Errorf("failed to save profit: %w", err)
	}
	return nil
}

func getUltimateValueProfit() float64 {
	data, _ := os.ReadFile(ultimateValueProfit)
	ultimateValueProfitText := string(data)
	profit, _ := strconv.ParseFloat(ultimateValueProfitText, 64)
	return profit
}

// RATIO SAVE
func writeUltimateValueRatio(ratio float64) error {
	ultimateValueRatioText := fmt.Sprint(ratio)
	err := os.WriteFile(ultimateValueRatio, []byte(ultimateValueRatioText), 0644)
	if err != nil {
		return fmt.Errorf("failed to save ratio: %w", err)
	}
	return nil
}

func getUltimateValueRatio() float64 {
	data, _ := os.ReadFile(ultimateValueRatio)
	ultimateValueRatioText := string(data)
	ratio, _ := strconv.ParseFloat(ultimateValueRatioText, 64)
	return ratio
}

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	outputText("Revenue: ")
	inputText(&revenue)
	if revenue <= 0 {
		fmt.Print("invalid value! revenue has to be greater than 0.")
		return
	}

	outputText("Expenses: ")
	inputText(&expenses)
	if expenses <= 0 {
		fmt.Print("invalid value! expenses has to be grater than 0.")
		return
	}

	outputText("Tax Rate in %: ")
	inputText(&taxRate)
	if taxRate <= 0 {
		fmt.Print("invalid value! expenses has to be grater than 0.")
		return
	}

	// variables def
	ebt, profit, ratio := calculateTax(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax - EBT: %0.2f\n", ebt)

	fmt.Printf("Earnings After Tax - profit: %0.2f\n", profit)

	fmt.Printf("Ratio: %0.2f\n", ratio)

	writeUltimateValueEbt(ebt)
	writeUltimateValueProfit(profit)
	writeUltimateValueRatio(ratio)

	outputText("The values saves in database are respectively: ")
	fmt.Print(getUltimateValueEbt(), " ", getUltimateValueProfit(), " ", getUltimateValueRatio())

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
