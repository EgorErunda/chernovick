package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFilePart(inputFilename, outFileName string, startPos int) error {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла %s: %w", inputFilename, err)
	}
	defer inputFile.Close()

	// Создаём или открываем файл для записи
	outputFile, err := os.Create(outFileName)
	if err != nil {
		return fmt.Errorf("ошибка создания файла %s: %w", outFileName, err)
	}
	defer outputFile.Close()

	// Перемещаемся к позиции startPos в исходном файле
	_, err = inputFile.Seek(int64(startPos), io.SeekStart)
	if err != nil {
		return fmt.Errorf("ошибка перемещения к позиции %d: %w", startPos, err)
	}

	// Копируем содержимое из inputFile в outputFile
	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("ошибка копирования данных: %w", err)
	}

	return nil
}
