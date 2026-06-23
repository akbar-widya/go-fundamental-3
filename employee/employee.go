package employee

import (
	"fmt"
	"strconv"
)

func findEmployee(id int) (string, bool) {
	employees := map[int]string{
		101: "Sarah",
		102: "John",
	}

	name, isEmployed := employees[id]

	return name, isEmployed
}

func Employee() {
	fmt.Println("--- 1. Converting Text (Value + Error) ---")

	userAge, err := strconv.Atoi("30")
	if err != nil {
		fmt.Println("Conversation failed:", err)
	} else {
		fmt.Println("Success! In 5 years, the user will be", userAge+5)
	}

	_, badErr := strconv.Atoi("thirty")
	if badErr != nil {
		fmt.Println("Caught a bad input:", badErr)
	}

	fmt.Println("\n--- 2. Custom functions (Value + Status) ---")

	employeeName, ok := findEmployee(101)
	if ok {
		fmt.Println("Employee found. Access granted for:", employeeName)
	}

	_, isFound := findEmployee(999)
	if !isFound {
		fmt.Println("Access denied: Employee ID 999 does not exist.")
	}
}