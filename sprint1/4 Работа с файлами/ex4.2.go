package main

import (
	"bufio"
	"fmt"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	file, err := os.Open(inputFilename)
	if err != nil {

		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0

	for scanner.Scan() {
		if lineCount == lineNum {
			return scanner.Text()
		}
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		return ""
	}

	return ""
}
