package helper

import (
	"fmt"
	"strings"
)

func PromptUser(prompt string) bool  {
	validResponse := "y"
	fmt.Print(prompt)

	var response string
	fmt.Scanln(&response)
	return strings.ToLower(response) == validResponse
}

func Retry() bool {
	return PromptUser("Try Again? (y/n): ")
}

func Home() bool {
	return PromptUser("Back to Home? (y/n): ")
}