package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main_() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}

//http.HandleFunc — функция, которая принимает путь URL и функцию-обработчика для пути с сигнатурой func(http.ResponseWriter, *http.Request). Эта функция будет вызываться для каждого запроса на указанный путь URL. Функция-обработчик может использовать http.ResponseWriter, чтобы записывать ответ, и *http.Request, чтобы получать информацию о запросе.
