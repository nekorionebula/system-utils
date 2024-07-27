package helper

import (
	"fmt"
	"strings"
)

func Retry() bool {
	fmt.Print("try again? (y/n): ")
	var try string
	fmt.Scanln(&try)
	return strings.ToLower(try) == "y"
}