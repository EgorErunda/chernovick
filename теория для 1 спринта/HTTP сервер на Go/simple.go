package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	http.ListenAndServe(":8080", nil)
}

//Этот код создаёт веб-сервер, который будет получать входящие запросы на порте 8080 и отправлять на них ответ "Hello World!".
