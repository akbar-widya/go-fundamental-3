package main

import (
	"errors"
	"fmt"
	"os"
)

func withdraw(balance int, amount int) (int, error) {
	if amount > balance {
		return balance, errors.New("insufficient funds")
	}
	return balance - amount, nil
}

func main() {
	fmt.Println("'--- 1. Banking Transaction (Value + Error) ---")

	currentBalance := 100

	newBalance, err := withdraw(currentBalance, 40)
	if err != nil {
		fmt.Println("Transaction failed:", err)
	} else {
		fmt.Println("Success! You withdrew $40. New balance is:", newBalance)
	}

	_, badErr := withdraw(currentBalance, 500)
	if badErr != nil {
		fmt.Println("Transaction blocked:", badErr)
	}

	fmt.Println("\n--- 2. Checking System Settings (Value + Status) ---")

	_, ok := os.LookupEnv("PATH")
	if ok {
		fmt.Println("System check: The 'PATH' variable exists on this computer.")
	}

	_, hasSecretKey := os.LookupEnv("MY_SECRET_API_KEY")
	if !hasSecretKey {
		fmt.Println("Security check: No secret key was found on this system.")
	}
}