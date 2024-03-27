package helpers

import (
	"bufio"
	"os"
)

// FormatInput removes all spaces from the input string.
func FormatInput(input string) string {
	output := []rune{}
	for _, r := range input {
		if r != ' ' {
			output = append(output, r)
		}
	}
	return string(output)
}

// GetInput returns the input string.
func GetInputOS() string {
	var input string
	os.Stdout.WriteString("Enter your calculation: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	return input
}
