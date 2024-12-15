package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseWriter struct {
	Name string `json:"name"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, `{"error":"параметр name обязателен"}`, http.StatusBadRequest)
		return
	}

	response := map[string]string{"name": name}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, `{"error":"ошибка при создании JSON"}`, http.StatusInternalServerError)
		return
	}

	log.Printf("%s\n", jsonResponse)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
