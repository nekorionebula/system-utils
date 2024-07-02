package kill

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func GetSleepDuration(durationStr string) time.Duration {
	for i := 0; i < 3; i++ {
		fmt.Print("The app will be terminated in (s): ")
		_, err := fmt.Scanf("%s\n", &durationStr)
		if err != nil {
			fmt.Println("Error :", err.Error())
			break
		}
		duration, parsErr := strconv.ParseFloat(durationStr, 64)
		if parsErr != nil {
			fmt.Println("invalid number format. please enter a valid number")
		} else {
			return time.Duration(duration * float64(time.Second))
		}
	}
	return 0
}

func getProcessID(processName string, duration time.Duration) (int, error) {
	//Run "tasklist" command
	cmd := exec.Command("tasklist")
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("output error %v", err.Error())
	}

	//Scan through the output line by line
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		//split the line into fields [" , ""]
		fields := strings.Fields(line)
		if len(fields) > 0 {
			//Check if the first field (process name) is an exact match
			if strings.EqualFold(fields[0], processName+".exe") {
				//convert the second field (PID) to an integer
				pid, pidErr := strconv.Atoi(fields[1])
				if pidErr != nil {
					return 0, fmt.Errorf("failed to retriece %v id", pidErr.Error())
				}
				//Sleep Duration before return
				fmt.Printf("%v will be terminated after %.2f minutes\n", processName, duration.Minutes())
				return pid, nil
			}
		}
	}
	if scanner.Err() != nil {
		return 0, fmt.Errorf("failed to scan %v", err)
	}
	return 0, fmt.Errorf("%v is not found", processName)
}

func kill(processName string, duration time.Duration) {
	pid, err := getProcessID(processName, duration)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	killNow := func() {
		time.Sleep(duration)
		cmd := exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(pid))
		if err = cmd.Run(); err != nil {
			fmt.Println("Error:", err.Error())
			return
		}
		fmt.Printf("%v has been kiled\n", processName)
	}
	go killNow()
}

func RunKill() {
	var (
		processName string
		durationStr string
	)
	fmt.Print("App name: ")
	_, err := fmt.Scanf("%s\n", &processName)
	if err != nil {
		fmt.Println("Error :", err.Error())
		return
	}
	//pid, err := getProcessID(processName)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//} else {
	//	fmt.Println("Process ID: ", pid)
	//}
	//parse
	duration := GetSleepDuration(durationStr)
	kill(processName, duration)
}
