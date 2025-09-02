package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)

	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var accountBalance float64 = getBalanceFromFile()

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
	
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
	writeBalanceToFile(balance)
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
	writeBalanceToFile(balance)
	return balance
}
