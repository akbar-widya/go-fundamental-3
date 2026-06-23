package main

import (
	"errors"
	"fmt"
)

func calculateTax(amount int) int {
	return (amount * 10) / 100
}

func formatCurrency(amount int) string {
	return fmt.Sprintf("%d.00", amount)
}

func processPayment(totalCost int, walletBalance int) (int, error) {
	if totalCost > walletBalance {
		return walletBalance, errors.New("payment declined: insufficient funds")
	}
	return walletBalance - totalCost, nil
}

func main() {
	fmt.Println("--- Checkout Process ---")

	itemPrice := 70
	wallet := 60

	tax := calculateTax(itemPrice)
	finalTotal := itemPrice + tax

	fmt.Println("Item Price:", formatCurrency(itemPrice))
	fmt.Println("Tax Added:", formatCurrency(tax))
	fmt.Println("Amount Due:", formatCurrency(finalTotal))

	fmt.Println("\n--- Processing Transaction ---")

	newBalance, err := processPayment(finalTotal, wallet)
	if err != nil {
		fmt.Println("System Alert:", err)
	} else {
		fmt.Println("Success! Thank you for your purchase.")
		fmt.Println("Remaining Wallet Balance:", formatCurrency(newBalance))
	}
}