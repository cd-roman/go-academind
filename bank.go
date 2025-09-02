package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000.00

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
	
		if choice == 1 {
			checkBalance(accountBalance)
		} else if choice == 2 {
			accountBalance = depositMoney(accountBalance)
		} else if choice == 3 {
			accountBalance = withdrawMoney(accountBalance)
		} else if choice == 4 {
			fmt.Println("Exiting...")
			break
		} else {
			fmt.Println("Invalid choice. Please try again.")
		}
	}

	fmt.Println("Thank you for using Go Bank!")

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
	return balance
}
