package main

import (
	"fmt"
	"errors"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("--- 1. Value and Error ---")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Failed:", err)
	} else {
		fmt.Println("Success! 10 / 2 =", result)
	}

	_, badErr := divide(10, 0)
	if badErr != nil {
		fmt.Println("Failure caught:", badErr)
	}

	fmt.Println("\n--- 2. The Comma-OK Idiom ---")

	inventory := map[string]int{"apples": 5}

	count, ok := inventory["apples"]
	if ok {
		fmt.Println("Item found! We have", count, "apples.")
	}

	fmt.Println("\n--- 3. The Blank Identifier (_) ---")

	_, hasOranges := inventory["oranges"]
	if !hasOranges {
		fmt.Println("Item missing: We do not sell oranges.")
	}
}