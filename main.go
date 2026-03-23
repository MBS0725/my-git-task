package main

import "fmt"

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorBlue   = "\033[34m"
)

func intToWord(n int) string {
	words := map[int]string{
		1: "one", 2: "two", 3: "three", 4: "four", 5: "five",
		6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten",
		11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen",
		15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen",
		19: "nineteen", 20: "twenty",
	}
	if n >= 1 && n <= 20 {
		return words[n]
	}
	return "big number"
}

func myCalculation(a int, b int, op string) string {
	var result int
	if op == "+" {
		result = a + b
	} else if op == "*" {
		result = a * b
	} else {
		return ColorRed + "Invalid operator" + ColorReset
	}
	return intToWord(a) + " " + op + " " + intToWord(b) + " = " + intToWord(result)
}

func main() {
	// Синий заголовок
	fmt.Println(ColorBlue + "=== Calculator for Code Review is ready! ===" + ColorReset)

	number := 3
	for i := 1; i <= 20; i++ {
		// Зеленый цвет для успешных расчетов
		res := myCalculation(i, number, "+")
		fmt.Println(ColorGreen + res + ColorReset)
	}
}