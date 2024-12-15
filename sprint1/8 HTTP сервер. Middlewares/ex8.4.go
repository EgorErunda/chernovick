package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

// Структура для ответа в формате JSON
type RPCResponse struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"` // Используем interface{} для пустого объекта
}

// Middleware Sanitize: проверяет имя и вызывает панику, если имя некорректно
func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(name) {
			panic("Invalid name") // Вызываем панику, если имя некорректно
		}
		next(w, r)
	}
}

// Middleware SetDefaultName: устанавливает имя по умолчанию, если оно не указано
func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "stranger"
		}
		r.URL.RawQuery = fmt.Sprintf("name=%s", name)
		next(w, r)
	}
}

// Middleware RPC: обрабатывает панику и возвращает JSON-ответ
func RPC(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Обрабатываем панику и возвращаем ошибку в формате JSON
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(RPCResponse{
					Status: "error",
					Result: struct{}{}, // Пустой объект
				})
			}
		}()

		// Вызываем следующий обработчик
		next(w, r)
	}
}

// Обработчик HelloHandler: возвращает ответ в формате JSON
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	response := RPCResponse{
		Status: "ok",
		Result: struct {
			Greetings string `json:"greetings"`
			Name      string `json:"name"`
		}{
			Greetings: "hello",
			Name:      name,
		},
	}

	// Устанавливаем заголовок Content-Type и возвращаем JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Устанавливаем обработчик с middleware
	http.HandleFunc("/", RPC(SetDefaultName(Sanitize(HelloHandler))))

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
