package kill

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type App struct {
	Name     string
	Duration time.Duration
}

func (app *App) getName() error {
	fmt.Print("App name: ")
	_, err := fmt.Scanln(&app.Name)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) getSleepDuration() error {
	for i := 0; i < 3; i++ {
		fmt.Print("The app will be terminated in (s): ")
		var durationStr string
		_, err := fmt.Scanln(&durationStr)
		if err != nil {
			return fmt.Errorf("error while scanning %w", err)
		}

		durationFl, err :=strconv.ParseFloat(durationStr, 32)
		if err == nil {
			app.Duration = time.Duration(durationFl * float64(time.Second))
			return nil
		}
		fmt.Println("invalid number format. please enter a valid number")
	}
	return fmt.Errorf("too many attempts")
}

func (app *App) scanApp() error {
	cmd := exec.Command("tasklist")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error getting tasklist output %w", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) > 0 && strings.EqualFold(fields[0], app.Name+".exe") {
			fmt.Printf("%s will be terminated after %.2f minutes\n", app.Name, app.Duration.Minutes())
			return nil
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}
	return fmt.Errorf("%s is not found", app.Name)
}



func  Kill() {
	app := &App{}
	if err := app.getName(); err != nil {
		fmt.Println("Error: cannot get the name ", err)
		return
	}

	if err := app.getSleepDuration(); err != nil {
		fmt.Println(err, "exiting...")
		time.Sleep(1 * time.Second)
		return
	}

	if err := app.scanApp(); err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(app.Duration)
	cmd := exec.Command("taskkill", "/F", "/T", "/IM", app.Name+".exe")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("runtime error:", err)
		fmt.Println("Output:", string(output))
	}
}