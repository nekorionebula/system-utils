package main

import (
	"context"
	"fmt"
	"time"

	"github.com/nekorionebula/system-utils/helper"
	"github.com/nekorionebula/system-utils/internal/kill"
	"github.com/nekorionebula/system-utils/scripts/termstyle"
)

func main() {
	var ctx context.Context
	var cancel context.CancelFunc
	for {
		fmt.Println("List of Utilities:")
		fmt.Printf("1. Kill App\n2. Shutdown\n0. Cancel All\n")
		fmt.Print("Task number: ")
		var input string
		_, err := fmt.Scanf("%s\n", &input)
		if err != nil {
			fmt.Println("Error input:", err)
		}
		switch input {
		case "0":
			if cancel != nil {
				cancel()
			}
			kill.CancelShutdown()
		case "1":
			ctx, cancel = context.WithCancel(context.Background())
			kill.KillApp(ctx)
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
		fmt.Println(termstyle.Yellow, "\"Warning: To ensure the process runs smoothly, do not close the program or terminal.\"", termstyle.Default)
		time.Sleep(500 * time.Millisecond)

		// if != n break
		if helper.Home() {
			break
		}
	}
	//cancel to avoid bugs
	if cancel != nil {
		cancel()
	}
	fmt.Print("Thank you :)")
}