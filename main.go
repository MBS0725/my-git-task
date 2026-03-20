package main

import (
	"fmt"
)


func intToWord(n int) string {
	// map — это словарь: число -> слово
	words := map[int]string{
		1: "one", 2: "two", 3: "three", 4: "four", 5: "five",
		6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten",
		11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen",
		15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen",
		19: "nineteen", 20: "twenty",
	}

	
	if n >= 1 && n <= 20 {
		return words[n] // возвращаем слово
	}

	// Если число больше 20
	return "big number"
}

// Функция выполняет математическую операцию и возвращает строку
func myCalculation(a int, b int, op string) string {
	var result int // переменная для хранения результата

	
	if op == "+" {
		result = a + b // сложение
	} else if op == "*" {
		result = a * b // умножение
	} else {
		// Если оператор неизвестный — ошибка
		panic("Invalid operator")
	}

	
	return intToWord(a) + " " + op + " " + intToWord(b) + " = " + intToWord(result)
}

func main() {
	// := — это короткое объявление переменной
	// number — это число, с которым будут выполняться операции
	number := 3

	// Цикл от 1 до 20
	for i := 1; i <= 20; i++ {
		// Вывод результата сложения
		fmt.Println(myCalculation(i, number, "+"))

		// Вывод результата умножения
		fmt.Println(myCalculation(i, number, "*"))
	}
}
	
