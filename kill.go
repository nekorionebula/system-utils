package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func getProcessID(processName string) (int, error) {
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
			if strings.EqualFold(fields[0], processName) {
				//convert the second field (PID) to an integer
				pid, err := strconv.Atoi(fields[1])
				if err != nil {
					return 0, fmt.Errorf("failed to retriece %v id", err.Error())
				}
				return pid, nil
			}
		}
	}
	if scanner.Err() != nil {
		return 0, fmt.Errorf("failed to scan %v", err.Error())
	}
	return 0, fmt.Errorf("%v is not found", processName)
}

func kill(processName string) {
	pid, err := getProcessID(processName)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	cmd := exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(pid))
	if err = cmd.Run(); err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Printf("%v has been kiled\n", processName)
}

func main() {
	processName := "spotify.exe"
	//pid, err := getProcessID(processName)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//} else {
	//	fmt.Println("Process ID: ", pid)
	//}
	time.Sleep(5 * time.Second)
	kill(processName)
}
