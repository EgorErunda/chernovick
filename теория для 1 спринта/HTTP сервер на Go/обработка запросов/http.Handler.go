package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	handler := &MyHandler{}
	http.ListenAndServe(":8080", handler)
}

//http.Handler — интерфейс, который определяет метод ServeHTTP и с его помощью принимает http.ResponseWriter и *http.Request. Объект, который реализует этот интерфейс, можно использовать для обработки HTTP-запросов. Например, http.FileServer реализует http.Handler, а эта функция обслуживает статические файлы из указанной директории.
