package main

import (
	"fmt"
	"net/http"
	"strings"
)

func StrangerHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "hello stranger")
		return
	}
	if !strings.ContainsAny(name, "abcdefghijklmnopqrstuvwxyzZYXWVUTSRQPONMLKJIHGFEDCBA") {
		fmt.Fprintf(w, "hello dirty hacker")
		return
	}
	fmt.Fprintf(w, "hello %s", name)
}

func main() {
	http.HandleFunc("/", StrangerHandler)
	http.ListenAndServe(":8080", nil)
}
