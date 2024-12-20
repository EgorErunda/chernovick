package main

import (
	"fmt"
	"net/http"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func main1() {
	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to the home page!")
	})
	// ...
}

//http.ServeMux — механизм маршрутизации. Он определяет, какие обработчики будут вызываться для каждого URL-адреса на сервере. Такой механизм полезен разработчикам веб-приложений, которые должны обрабатывать множество запросов на разные URL-адреса.
