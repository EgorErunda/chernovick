package main

import (
	"fmt"
	"io"
)

// Copy копирует n байт из r в w.
// Если доступно меньше n байт, копирует все доступные данные.
// Если возникает ошибка, возвращает её.
func Copy(r io.Reader, w io.Writer, n uint) error {
	// Создаём буфер для чтения данных
	buf := make([]byte, n)

	// Читаем данные из r
	bytesRead, err := r.Read(buf)
	if err != nil && err != io.EOF {
		return fmt.Errorf("ошибка чтения: %w", err)
	}

	// Если прочитали меньше n байт, обрезаем буфер
	if uint(bytesRead) < n {
		buf = buf[:bytesRead]
	}

	// Записываем данные в w
	bytesWritten, err := w.Write(buf)
	if err != nil {
		return fmt.Errorf("ошибка записи: %w", err)
	}

	// Проверяем, что все данные были записаны
	if bytesWritten != bytesRead {
		return fmt.Errorf("не удалось записать все данные: записано %d байт, ожидалось %d байт", bytesWritten, bytesRead)
	}

	return nil
}
