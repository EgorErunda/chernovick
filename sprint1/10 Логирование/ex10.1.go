package main

import (
	"fmt"
	"os"
)

func WriteToLogFile(message string, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(message + "\n")
	if err != nil {
		return fmt.Errorf("не удалось записать в файл: %w", err)
	}

	return nil
}

// func main() {
// 	// Пример использования функции
// 	fileName := "output.txt"
// 	message := "hello world"

// 	err := WriteToLogFile(message, fileName)
// 	if err != nil {
// 		fmt.Printf("Ошибка: %v\n", err)
// 		return
// 	}

// 	fmt.Println("Сообщение успешно записано в файл.")
// }
