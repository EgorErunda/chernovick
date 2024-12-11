package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// run считывает параметры из командной строки и записывает их в файл config.txt.
func run() error {
	// Проверяем, что передано ровно 3 аргумента
	if len(os.Args) != 4 {
		return errors.New("использование: main <высота> <ширина> <процент заполнения>")
	}

	// Считываем аргументы
	height, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return errors.New("высота должна быть целым числом")
	}

	width, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return errors.New("ширина должна быть целым числом")
	}

	fillPercent, err := strconv.Atoi(os.Args[3])
	if err != nil {
		return errors.New("процент заполнения должен быть целым числом")
	}

	// Проверяем, что процент заполнения находится в диапазоне от 0 до 100
	if fillPercent < 0 || fillPercent > 100 {
		return errors.New("процент заполнения должен быть от 0 до 100")
	}

	// Формируем строку для записи в файл
	configString := fmt.Sprintf("%dx%d %d%%", height, width, fillPercent)

	// Записываем строку в файл config.txt
	err = os.WriteFile("config.txt", []byte(configString), 0644)
	if err != nil {
		return fmt.Errorf("не удалось записать в файл: %w", err)
	}

	return nil
}
