package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("Failed to find file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return defaultValue, errors.New("Failed to parse stored value")
	}

	return value, nil
}

func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func main() {
	var accountBalance, err = getFloatFromFile(accountBalanceFile, 0)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------")
		// panic("Failed to get account balance")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			checkBalance(accountBalance)
		case 2:
			accountBalance = depositMoney(accountBalance)
		case 3:
			accountBalance = withdrawMoney(accountBalance)
		case 4:
			fmt.Println("Exiting...")
			fmt.Println("Thank you for using Go Bank!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func checkBalance(balance float64) {
	fmt.Printf("Your current balance is: %.2f\n", balance)
}

func depositMoney(balance float64) float64 {
	var amount float64
	fmt.Print("Enter amount to deposit: ")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Invalid deposit amount. Must be greater than 0.")
		return balance
	}

	balance += amount
	fmt.Printf("You have deposited: %.2f\n", amount)
	fmt.Printf("Your new balance is: %.2f\n", balance)
	writeFloatToFile(balance, accountBalanceFile)
	return balance
}

func withdrawMoney(balance float64) float64 {
	var amount float64
	fmt.Print("Enter amount to withdraw: ")
	fmt.Scan(&amount)

	if amount > balance {
		fmt.Println("Insufficient funds.")
		return balance
	} else {
		balance -= amount
		fmt.Printf("You have withdrawn: %.2f\n", amount)
	}

	fmt.Printf("Your new balance is: %.2f\n", balance)
	writeFloatToFile(balance, accountBalanceFile)
	return balance
}


