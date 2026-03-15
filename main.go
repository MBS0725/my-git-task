package main

import (
	"fmt"
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
		panic("Invalid operator")
	}

	return intToWord(a) + " " + op + " " + intToWord(b) + " = " + intToWord(result)
}

func main() {
	number := 3

	for i := 1; i <= 20; i++ {
		fmt.Println(myCalculation(i, number, "+"))
		fmt.Println(myCalculation(i, number, "*"))
	}
}

