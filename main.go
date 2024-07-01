package main

import (
	"fmt"
	"github.com/nekorionebula/system-utils/internal/kill"
)

var (
	input string
	try   string
)

func Input() {
	fmt.Println("List of Utilities:")
	fmt.Println("1. Kill App")
	_, err := fmt.Scanf("%s\n", &input)
	if err != nil {
		fmt.Println("Error:", err)
	}
	switch input {
	case "1":
		kill.RunKill()
	default:
		fmt.Println("Input correct number.")
		fmt.Println("Try again? (y/n): ")
		_, _ = fmt.Scanf("%s\n", &try)
		switch try {
		case "y", "Y", "yes", "Yes", "YES":
			fmt.Println()
			Input()
		case "n", "N", "no", "No", "NO":
			return
		}
	}
}

func main() {
	Input()
}
