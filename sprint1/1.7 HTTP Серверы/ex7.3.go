package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// Структура для хранения текущего состояния чисел Фибоначчи
type FibonacciState struct {
	mu   sync.Mutex
	a, b int
}

// Метод для получения следующего числа Фибоначчи
func (fs *FibonacciState) Next() int {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	// Вычисляем следующее число Фибоначчи
	next := fs.a
	fs.a, fs.b = fs.b, fs.a+fs.b
	return next
}

// Глобальная переменная для хранения состояния
var fibState = FibonacciState{a: 0, b: 1}

// Счётчик запросов
var requestCount uint64

// Обработчик запросов на получение числа Фибоначчи
func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	// Увеличиваем счётчик запросов
	atomic.AddUint64(&requestCount, 1)

	// Получаем следующее число Фибоначчи
	next := fibState.Next()

	// Возвращаем число в ответе
	fmt.Fprintf(w, "%d", next)
}

// Обработчик для получения метрик
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	// Формируем строку с метриками
	metrics := "rpc_duration_milliseconds_count " + strconv.FormatUint(requestCount, 10)

	// Возвращаем метрики в ответе
	fmt.Fprint(w, metrics)
}

// Функция для запуска сервера
func StartServer(timeout time.Duration) {
	// Устанавливаем обработчики для маршрутов

	//При запросе к корневому маршруту (/) сервер будет вызывать функцию FibonacciHandler, которая возвращает следующее число Фибоначчи.
	http.HandleFunc("/", FibonacciHandler)
	//При запросе к маршруту (/metrics) сервер будет вызывать функцию MetricsHandler, которая возвращает метрики.
	http.HandleFunc("/metrics", MetricsHandler)

	// Запускаем сервер на порту :8080
	fmt.Println("Сервер запущен на http://localhost:8080")
	server := &http.Server{ //Создаётся объект http.Server с настройками:
		Addr:         ":8080",
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}
	err := server.ListenAndServe() //Запускает сервер.
	if err != nil {
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
	}
}
