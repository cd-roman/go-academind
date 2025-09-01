package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)
	
	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	formattedFutureValue := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFutureRealValue := fmt.Sprintf("Future Real Value: %.2f\n", futureRealValue)

	fmt.Print(formattedFutureValue)
	fmt.Print(formattedFutureRealValue)
}