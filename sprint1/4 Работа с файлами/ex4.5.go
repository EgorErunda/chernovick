package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	file, err := os.Open(inputFileName)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла %s: %w", inputFileName, err)
	}
	defer file.Close()

	// Создаём буферизованный читатель
	scanner := bufio.NewScanner(file)
	// Результат — массив строк, соответствующих диапазону
	var result []string

	// Формат даты в логе
	dateFormat := "02.01.2006"

	// Читаем файл построчно
	for scanner.Scan() {
		line := scanner.Text()

		// Извлекаем дату из строки
		dateStr := strings.SplitN(line, " ", 2)[0]
		logDate, err := time.Parse(dateFormat, dateStr)
		if err != nil {
			return nil, fmt.Errorf("ошибка парсинга даты в строке '%s': %w", line, err)
		}

		// Проверяем, попадает ли дата в диапазон [start..end]
		if logDate.After(end) || logDate.Before(start) {
			continue // Пропускаем строки, не попадающие в диапазон
		}

		// Добавляем строку в результат
		result = append(result, line)
	}

	// Проверяем, была ли ошибка при сканировании
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	// Если ни одна строка не попала в диапазон, возвращаем ошибку
	if len(result) == 0 {
		return nil, errors.New("ни одна строка не попала в указанный диапазон")
	}

	return result, nil
}
