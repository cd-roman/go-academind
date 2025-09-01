package main

import (
	"fmt"
)

func main() {
	revenue := getUserInput("Enter your revenue: ")
	expenses := getUserInput("Enter your expenses: ")
	taxRate := getUserInput("Enter your tax rate: ")

	earningsBeforeTax, profit, earningsRatio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Your earnings before tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Your profit after tax: %.2f\n", profit)
	fmt.Printf("Your earnings ratio: %.2f\n", earningsRatio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	earningsRatio := earningsBeforeTax / profit

	return earningsBeforeTax, profit, earningsRatio
}