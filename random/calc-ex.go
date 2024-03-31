package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	calculate()
}

func getCharString() string {
	reader := bufio.NewReader(os.Stdin)

	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(char)
}

func calculate() {
	var charString string

	var expression string

	var num1 float64
	var num2 float64
	var result float64

	var display string

	operator := ""
	temp := ""
	result = 0.0

	for charString != "q" {
		fmt.Println("Enter a character:")
		charString = getCharString()

		switch charString {
		case "c":
			expression = ""
			num1 = 0
			num2 = 0
			result = 0
			operator = ""
			temp = ""
			display = ""
		case "=":
			// evaluate expression
			switch operator {
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			case "*":
				result = num1 * num2
			case "/":
				result = num1 / num2
			}
			// set num1 to result
			num1 = result
			// set num2 to 0
			num2 = 0
			temp = ""
			// set operator to ""
			operator = "="
			display = strconv.FormatFloat(result, 'f', -1, 64)
			expression = ""

		case "+", "-", "*", "/":
			// set operator
			if operator != "" {
				// set operator to charString
				operator = charString
				expression += charString
				// evaluate expression
				switch operator {
				case "+":
					result = num1 + num2
				case "-":
					result = num1 - num2
				case "*":
					result = num1 * num2
				case "/":
					result = num1 / num2
				}
				// set num1 to result
				num1 = result
				// set num2 to 0
				num2 = 0
				temp = ""
				display = strconv.FormatFloat(result, 'f', -1, 64) + charString
			} else {
				operator = charString
				num1, _ = strconv.ParseFloat(expression, 64)
				num2 = 0
				temp = ""
				expression += charString
				display = expression
			}
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			// append to expression string
			temp += charString
			expression += charString
			if operator == "" {
				num1, _ = strconv.ParseFloat(temp, 64)
				display = strconv.FormatFloat(num1, 'f', -1, 64)
			} else {
				num2, _ = strconv.ParseFloat(temp, 64)
				display = strconv.FormatFloat(num2, 'f', -1, 64)
			}

		default:
			// invalid input
			fmt.Println("Invalid input")
		}

		fmt.Println("Character:", charString)
		fmt.Println("Display:", display)
		fmt.Println("Expression:", expression)
		fmt.Println("Operator:", operator)
		fmt.Println("Num1:", num1)
		fmt.Println("Num2:", num2)
		fmt.Println("Result:", result)
	}
}
