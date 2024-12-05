package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например, \"hello world\" + \"from go\" или \"repeat this\" * 3):")

	// Читаем ввод полностью, включая пробелы
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	// Убираем перенос строки
	input = strings.TrimSpace(input)

	// Выполняем вычисление
	result, err := calculate(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Ограничиваем длину результата до 40 символов
	if len(result) > 40 {
		result = result[:40] + "..."
	}
	fmt.Println("Результат:", result)
}

func calculate(input string) (string, error) {
	input = strings.TrimSpace(input)

	// Парсим выражение
	left, operator, right, err := parseExpression(input)
	if err != nil {
		return "", err
	}

	if len(left) > 10 {
		return "", errors.New("длина строки должна быть не более 10 символов")
	}

	switch operator {
	case "+":
		if len(right) > 10 {
			return "", errors.New("длина строки должна быть не более 10 символов")
		}
		return left + right, nil

	case "-":
		return strings.ReplaceAll(left, right, ""), nil

	case "*":
		n, err := strconv.Atoi(right)
		if err != nil || n < 1 || n > 10 {
			return "", errors.New("множитель должен быть целым числом от 1 до 10")
		}
		return strings.Repeat(left, n), nil

	case "/":
		n, err := strconv.Atoi(right)
		if err != nil || n < 1 || n > 10 {
			return "", errors.New("делитель должен быть целым числом от 1 до 10")
		}
		if len(left)%n != 0 {
			return "", errors.New("строка не делится нацело на указанное число")
		}
		return left[:len(left)/n], nil

	default:
		return "", errors.New("неизвестная операция")
	}
}

// Парсинг выражения с учётом двойных кавычек
func parseExpression(input string) (string, string, string, error) {
	// Определяем левую строку
	left, rest, err := extractQuotedPart(input)
	if err != nil {
		return "", "", "", err
	}

	// Ищем оператор
	rest = strings.TrimSpace(rest)
	var operator string
	if strings.HasPrefix(rest, "+") {
		operator = "+"
	} else if strings.HasPrefix(rest, "-") {
		operator = "-"
	} else if strings.HasPrefix(rest, "*") {
		operator = "*"
	} else if strings.HasPrefix(rest, "/") {
		operator = "/"
	} else {
		return "", "", "", errors.New("оператор отсутствует или неизвестен")
	}

	// Убираем оператор и пробелы
	rest = strings.TrimSpace(rest[len(operator):])

	// Определяем правую часть
	right, remaining, err := extractQuotedPart(rest)
	if err != nil {
		// Если правая часть не в кавычках, проверяем на число
		right = rest
		if _, err := strconv.Atoi(right); err != nil {
			return "", "", "", errors.New("правая часть должна быть числом или строкой в двойных кавычках")
		}
		return left, operator, right, nil
	}

	if len(strings.TrimSpace(remaining)) != 0 {
		return "", "", "", errors.New("некорректный формат выражения")
	}

	return left, operator, right, nil
}

// Извлекает содержимое в двойных кавычках и возвращает оставшуюся часть строки
func extractQuotedPart(input string) (string, string, error) {
	input = strings.TrimSpace(input)

	if len(input) == 0 || input[0] != '"' {
		return "", input, errors.New("элемент должен начинаться с двойной кавычки")
	}

	end := strings.Index(input[1:], "\"")
	if end == -1 {
		return "", input, errors.New("строка должна быть завершена двойной кавычкой")
	}
	end += 1 // Учитываем смещение из-за кавычки

	content := input[1:end]
	remaining := input[end+1:]
	return content, remaining, nil
}
