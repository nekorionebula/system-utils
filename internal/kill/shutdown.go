package kill

import (
	"fmt"
	"os/exec"
	"strconv"
)

func Shutdown() {
	var (
		timeStr  string
		time     int
		errParse error
	)

	const maxAttempt = 3

	for i := 0; i < maxAttempt; i++ {

		fmt.Print("Enter time in seconds: ")
		_, _ = fmt.Scan(&timeStr)
		time, errParse = strconv.Atoi(timeStr)

		//error
		if errParse != nil {
			fmt.Println("Error: please enter a valid time in second(s)")
			if i < maxAttempt-1 {
				var retry string
				fmt.Print("Do you want to retry? (y/n): ")
				_, _ = fmt.Scan(&retry)

				if retry != "y" {
					fmt.Println("Exiting...")
					return
				}
			} else {
				fmt.Println("Maximum attempted reached. Exiting...")
				return
			}
		} else {
			break
		}
	}

	command := fmt.Sprintf("shutdown -s -t %v", time)
	cmd := exec.Command("powershell", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Program executed successfully", string(output))
}
