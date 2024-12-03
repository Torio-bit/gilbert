package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Text struct {
	Content string
}

func (t *Text) stringModification() {

	fmt.Println("Удаляет лишних пробелов и да по другому не как нужно оби эти команды") /* Подробнее в обсидиане в разделе библиотека стринг */

	fmt.Println(strings.Join(strings.Fields(t.Content), " "))

	fmt.Println("Он показывает что такого слово есть или нет в строке (Go)")

	fmt.Println(strings.Contains(t.Content, "Go"))

	fmt.Println(strings.Replace(t.Content, "Go", "Golang", -1))

	fmt.Println(strings.ReplaceAll(t.Content, "GO", "Golang"))
}

func main() {

	input := " Егор ГОРп а    ываыоа    "

	// Разделение строки на подстроки
	words := strings.Fields(input)

	// Вывод результата
	fmt.Println(words)

	text := &Text{}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите типо строку")

	for scanner.Scan() {
		text.Content = scanner.Text()
		text.stringModification()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}
