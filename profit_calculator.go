// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var revenue, expenses, taxRate float64

// 	fmt.Print("Enter your revenue: ")
// 	fmt.Scan(&revenue)

// 	fmt.Print("Enter your expenses: ")
// 	fmt.Scan(&expenses)

// 	fmt.Print("Enter your tax rate: ")
// 	fmt.Scan(&taxRate)

// 	earningsBeforeTax := revenue - expenses
// 	profit := earningsBeforeTax * (1 - taxRate/100)
// 	earningsRatio := earningsBeforeTax / profit

// 	fmt.Println("Your earnings before tax:", earningsBeforeTax)
// 	fmt.Println("Your profit after tax:", profit)
// 	fmt.Println("Your earnings ratio:", earningsRatio)
// }