package main

import (
	"bufio"
	"fmt"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	// Открываем файл
	file, err := os.Open(inputFilename)
	if err != nil {
		// Если файл не существует или произошла ошибка, возвращаем пустую строку
		return ""
	}
	defer file.Close() // Закрываем файл после завершения работы

	// Создаём буферизованный читатель
	scanner := bufio.NewScanner(file)

	// Счётчик строк
	lineCount := 0

	// Читаем файл построчно
	for scanner.Scan() {
		lineCount++
		if lineCount == lineNum {
			return scanner.Text()
		}
	}

	// Если произошла ошибка при сканировании, выводим её
	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		return ""
	}

	// Если строка с указанным номером не найдена, возвращаем пустую строку
	return ""
}
