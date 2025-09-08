// package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	revenue := getUserInput("Enter your revenue: ")
// 	expenses := getUserInput("Enter your expenses: ")
// 	taxRate := getUserInput("Enter your tax rate: ")

// 	earningsBeforeTax, profit, earningsRatio := calculateFinancials(revenue, expenses, taxRate)

// 	// Display results on console
// 	fmt.Printf("Your earnings before tax: %.2f\n", earningsBeforeTax)
// 	fmt.Printf("Your profit after tax: %.2f\n", profit)
// 	fmt.Printf("Your earnings ratio: %.2f\n", earningsRatio)

// 	// Store results to file
// 	err := storeResultsToFile(earningsBeforeTax, profit, earningsRatio)
// 	if err != nil {
// 		fmt.Printf("Warning: Could not save results to file: %v\n", err)
// 	} else {
// 		fmt.Println("Results saved to 'profit_results.txt'")
// 	}
// }

// func storeResultsToFile(earningsBeforeTax, profit, earningsRatio float64) error {
// 	results := fmt.Sprintf("Earnings Before Tax: %.2f\nProfit: %.2f\nRatio: %.2f\n", earningsBeforeTax, profit, earningsRatio)
// 	return os.WriteFile("profit_results.txt", []byte(results), 0644)
// }

// func validateUserInput(value float64) error {
// 	if value <= 0 {
// 		return errors.New("value must be greater than 0")
// 	}
// 	return nil
// }

// func getUserInput(infoText string) float64 {
// 	var userInput float64
// 	fmt.Print(infoText)
// 	fmt.Scan(&userInput)
	
// 	err := validateUserInput(userInput)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		os.Exit(1)
// 	}
	
// 	return userInput
// }

// func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
// 	earningsBeforeTax := revenue - expenses
// 	profit := earningsBeforeTax * (1 - taxRate/100)
// 	earningsRatio := earningsBeforeTax / profit

// 	return earningsBeforeTax, profit, earningsRatio
// }