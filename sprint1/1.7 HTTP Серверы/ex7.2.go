package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Fibonacci struct {
	mu   sync.Mutex
	x, y int
}

// Метод для получения следующего числа фибоначчи
func (f *Fibonacci) Next() int {
	f.mu.Lock()
	defer f.mu.Unlock()

	next := f.x
	f.x, f.y = f.y, f.x+f.y
	return next
}

// глобальная переменная для хранения состояния
var FibState = Fibonacci{x: 0, y: 1}

// обработчик запросов
func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	//получаем следующее число
	next := FibState.Next()
	fmt.Fprintf(w, "%d", next)
}

// Функция для запуска сервера
func startServer(timeout time.Duration) {
	// Устанавливаем обработчик для маршрута "/"
	http.HandleFunc("/", fibonacciHandler)

	// Запускаем сервер на порту :8080
	fmt.Println("Сервер запущен на http://localhost:8080")
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
	}
}
