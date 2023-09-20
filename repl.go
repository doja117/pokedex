package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := CleanInput(text)
		fmt.Println("echo ", cleaned)
	}
}

func CleanInput(s string) []string {
	lowered := strings.ToLower(s)
	words := strings.Fields(lowered)
	return words
}
