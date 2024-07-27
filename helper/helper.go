package helper

import (
	"fmt"
	"strings"
)

func PromptUser(prompt string, validResponse string) bool  {
	fmt.Print(prompt)

	var response string
	fmt.Scanln(&response)
	return strings.ToLower(response) == validResponse
}

func Retry() bool {
	return PromptUser("Try Again? (y/n): ", "y")
}

func Home() bool {
	return PromptUser("Back to Home? (y/n): ", "n")
}