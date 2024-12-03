package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Введите выражение (например: \"hello\" + \"world\", \"hello\" - \"lo\", \"hello\" * 3, \"hello\" / 2):")
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic("Ошибка при вводе. Убедитесь, что строка введена правильно.")
	}

	// Удаляем пробелы по краям
	input = strings.TrimSpace(input)

	// Проверяем на наличие операций
	var firstStr, secondStr string
	var operation string
	var num int

	// Проверка на сложение
	if strings.Contains(input, "+") {
		parts := strings.Split(input, "+")
		if len(parts) != 2 {
			panic("Некорректное выражение. Используйте формат: \"строка1\" + \"строка2\".")
		}
		firstStr = strings.TrimSpace(parts[0])
		secondStr = strings.TrimSpace(parts[1])
		operation = "+"
	} else if strings.Contains(input, "-") { // Проверка на вычитание
		parts := strings.Split(input, "-")
		if len(parts) != 2 {
			panic("Некорректное выражение. Используйте формат: \"строка1\" - \"строка2\".")
		}
		firstStr = strings.TrimSpace(parts[0])
		secondStr = strings.TrimSpace(parts[1])
		operation = "-"
	} else if strings.Contains(input, "*") { // Проверка на умножение
		parts := strings.Split(input, "*")
		if len(parts) != 2 {
			panic("Некорректное выражение. Используйте формат: \"строка\" * число.")
		}
		firstStr = strings.TrimSpace(parts[0])
		num = parseNumber(strings.TrimSpace(parts[1]))
		operation = "*"
	} else if strings.Contains(input, "/") { // Проверка на деление
		parts := strings.Split(input, "/")
		if len(parts) != 2 {
			panic("Некорректное выражение. Используйте формат: \"строка\" / число.")
		}
		firstStr = strings.TrimSpace(parts[0])
		num = parseNumber(strings.TrimSpace(parts[1]))
		operation = "/"
	} else {
		panic("Недопустимая операция. Используйте одну из следующих: +, -, *, /.")
	}

	// Выполняем операцию
	result := ""
	switch operation {
	case "+":
		result = firstStr + secondStr
	case "-":
		result = strings.ReplaceAll(firstStr, secondStr, "")
		if result == "" {
			result = firstStr
		}
	case "*":
		result = repeatString(firstStr, num)
	case "/":
		result = divideString(firstStr, num)
	}

	// Ограничение длины результата
	if len(result) > 40 {
		result = result[:40] + "..."
	}

	fmt.Println("Результат:", result)
}

// Функция для преобразования строки в число
func parseNumber(s string) int {
	var num int
	_, err := fmt.Sscanf(s, "%d", &num)
	if err != nil || num < 1 || num > 10 {
		panic("Число должно быть целым от 1 до 10.")
	}
	return num
}

// Функция для повторения строки
func repeatString(str string, count int) string {
	var builder strings.Builder
	for i := 0; i < count; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

// Функция для деления строки на число
func divideString(str string, count int) string {
	if count <= 0 {
		panic("Деление на ноль или меньше недопустимо.")
	}
	length := len(str)
	if length < count {
		return "" // Результат при делении строки на число больше ее длины
	}
	partLength := length / count
	return str[:partLength]
}
