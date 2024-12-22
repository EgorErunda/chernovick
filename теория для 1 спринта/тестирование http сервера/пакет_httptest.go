//NewRequest имитирует запрос, который будет использоваться для обслуживания обработчика
//NewRecorder заменяет ResponseWriter: он обрабатывает HTTP-ответ и сравнивает его с ожидаемым результатом

package main

import (
	"fmt"
	"io"

	"net/http"

	"net/http/httptest"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	//создаёт новый макет запроса для вашего сервера. Вы передаёте его своему обработчику в качестве объекта запроса
	req := httptest.NewRequest("GET", "http://google.com", nil)
	//w := httptest.NewRecorder() записывает все ответы от обработчика
	w := httptest.NewRecorder()
	//handler(w, req) вызывает функцию обработчика
	handler(w, req)
	//resp := w.Result() возвращает результат http-сервера
	resp := w.Result()
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("body: ", string(body))
	fmt.Println("status code: ", resp.StatusCode)
}
