package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 3.5
	var investmentAmount float64
	var expertedReturnRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Experted Return Rate: ")
	fmt.Scan(&expertedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expertedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future Value: %0.2f\n", futureValue)

	fmt.Printf("Future Real Value: %0.2f\n", futureRealValue)
}
