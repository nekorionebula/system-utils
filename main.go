package main

import (
	"fmt"
	"github.com/nekorionebula/system-utils/internal/kill"
	"github.com/nekorionebula/system-utils/scripts/termstyle"
)

func main() {
	Input()
}

var (
	input    string
	try      string
	backHome string
)

func Input() {
	for {
		fmt.Println("List of Utilities:")
		fmt.Printf("1. Kill App\n2. Shutdown\n")
		fmt.Print("Task number: ")
		//delete: catch err
		_, _ = fmt.Scanf("%s\n", &input)
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

		//toGoHome
		fmt.Println(termstyle.Yellow, "\"Warning: If you want to kill the app, don't close the program or terminal\"", termstyle.Default)
		fmt.Print("Back to home? (y/n): ")
		_, _ = fmt.Scan(&backHome)
		fmt.Println()
		if backHome != "y" {
			break
		}
	}
	fmt.Println("Thank you :)")
}
