package main

import (
	"strings"
	"unicode"
)

// UpperWriter реализует интерфейс io.Writer
type UpperWriter struct {
	UpperString string // Поле для хранения строки в верхнем регистре
}

// Write реализует метод интерфейса io.Writer
func (w UpperWriter) Write(p []byte) (n int, err error) {
	// Преобразуем байты в строку
	str := string(p)

	// Переводим строку в верхний регистр
	upperStr := strings.Map(func(r rune) rune {
		return unicode.ToUpper(r)
	}, str)

	// Записываем строку в поле UpperString
	w.UpperString += upperStr

	// Возвращаем количество записанных байтов и nil ошибку
	return len(p), nil
}
