package kill

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/nekorionebula/system-utils/helper"
)
func CancelShutdown(){
	exec.Command("shutdown", "/a").Run()
}

func getDuration() int{
	for i := 0; i < 3; i++ {

		fmt.Print("Enter time in seconds: ")
		var durationStr string
		_, _ = fmt.Scanln(&durationStr)
		duration, err := strconv.Atoi(durationStr)

		//error
		if err != nil || duration < 0 {
			fmt.Println("Error: please enter a valid time of seconds")
			if helper.Retry() {
				continue
			}
			fmt.Println("Exiting...")
			return 0
		}
		return duration	
	}
	fmt.Println("Maximum attempts reached Exiting...")
	return 0
}

func System() { 
	duration := getDuration()
	if duration > 0 && duration <= 30 {
		fmt.Println("Less than 30 seconds. Attempting to cancel shutdown...")
		return
	}
	command := fmt.Sprintf("shutdown -s -t %d", duration)
	cmd := exec.Command("powershell", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("success", string(output))
}