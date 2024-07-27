package main

import (
	"fmt"

	"github.com/nekorionebula/system-utils/helper"
	"github.com/nekorionebula/system-utils/internal/kill"
	"github.com/nekorionebula/system-utils/scripts/termstyle"
)

func main() {
	for {
		fmt.Println("List of Utilities:")
		fmt.Printf("1. Kill App\n2. Shutdown\n0. Exit\n")
		fmt.Print("Task number: ")
		//delete: catch err
		var input string
		_, err := fmt.Scanf("%s\n", &input)
		if err != nil {
			fmt.Println("Error input:", err)
		}
		switch input {
		case "0":
			return
		case "1":
			kill.KillApp()	
		case "2":
			kill.System()
		default:
			fmt.Println("Input correct number.")
			if helper.Retry() {
				continue
			}
			return	
		}
		//toGoHome
		fmt.Println(termstyle.Yellow, "\"Warning: If you want run the process, don't close the program or terminal\"", termstyle.Default)
			if !helper.Home() {
				break
			}
	}
	fmt.Print("Thank you :)")
}