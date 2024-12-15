package main

import (
	"bufio"
	"bytes"
	"io"
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	br := bufio.NewReader(r) // Буферизованный читатель для улучшения производительности(позволяет читаь не по одному байту,а блоками)

	// Создаём буфер, размер которого кратен длине последовательности
	buf := make([]byte, 4096) // Размер буфера 4096 байт (можно настроить)

	for {
		// Читаем данные в буфер
		n, err := br.Read(buf)
		if err != nil && err != io.EOF {
			return false, err // Возвращаем ошибку, если она не io.EOF
		}

		// Ищем последовательность в буфере
		if bytes.Contains(buf[:n], seq) {
			return true, nil // Последовательность найдена
		}

		// Если достигнут конец данных, завершаем поиск
		if err == io.EOF {
			break
		}
	}

	return false, nil // Последовательность не найдена
}
