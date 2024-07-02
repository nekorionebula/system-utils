package main

import (
	"fmt"
	"github.com/nekorionebula/system-utils/internal/kill"
)

func main() {
	Input()
}

var (
	input string
	try   string
)

func Input() {
	fmt.Println("List of Utilities:")
	fmt.Printf("1. Kill App\n2. Shutdown\n")
	fmt.Print("Task number: ")
	_, err := fmt.Scanf("%s\n", &input)
	if err != nil {
		fmt.Println("Error:", err)
	}
	switch input {
	case "1":
		kill.RunKill()
	case "2":
		kill.Shutdown()
	default:
		fmt.Println("Input correct number.")
		fmt.Println("Try again? (y/n): ")
		_, _ = fmt.Scanf("%s\n", &try)

		if try != "y" {
			fmt.Println("Exiting...")
			return
		} else {
			Input()
		}
	}
}
