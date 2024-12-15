package main

import (
	"fmt"
	"os"
)

func ModifyFile(filename string, pos int, val string) error {
	// Открываем файл на чтение и запись
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла %s: %w", filename, err)
	}
	defer file.Close() // Закрываем файл после завершения работы

	// Перемещаемся к позиции pos в файле
	_, err = file.Seek(int64(pos), os.SEEK_SET)
	if err != nil {
		return fmt.Errorf("ошибка перемещения к позиции %d: %w", pos, err)
	}

	// Записываем значение val в файл, начиная с позиции pos
	_, err = file.WriteString(val)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}

	return nil
}
