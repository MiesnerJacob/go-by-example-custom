package main

import (
	"fmt"
	"strings"
)

func prefixGenerator(prefix string) func(string) string {
	return func(text string) string {
		return prefix + ": " + strings.TrimSpace(text)
	}
}

func main() {
	logError := prefixGenerator("ERROR")
	logInfo := prefixGenerator("INFO")

	fmt.Println(logError("Something went wrong"))
	fmt.Println(logInfo("Operation successful"))
}
