package main

// import (
// 	"fmt"
// 	"net/http"
// 	"regexp"
// )

// func Sanitize(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		name := r.URL.Query().Get("name")
// 		if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(name) {
// 			name = "dirty hacker"
// 		}
// 		newURL := *r.URL
// 		newURL.RawQuery = fmt.Sprintf("name=%s", name)
// 		newReq := r.Clone(r.Context())
// 		newReq.URL = &newURL

// 		next.ServeHTTP(w, newReq)
// 	}
// }

// func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		name := r.URL.Query().Get("name")
// 		if name == "" {
// 			name = "stranger"
// 		}
// 		newURL := *r.URL
// 		newURL.RawQuery = fmt.Sprintf("name=%s", name)
// 		newReq := r.Clone(r.Context())
// 		newReq.URL = &newURL

// 		next.ServeHTTP(w, newReq)
// 	}
// }

// func HelloHandler(w http.ResponseWriter, r *http.Request) {
// 	name := r.URL.Query().Get("name")
// 	fmt.Fprintf(w, "Hello, %s!", name)
// }

// // func main() {
// // 	http.HandleFunc("/", SetDefaultName(Sanitize(HelloHandler)))

// // 	fmt.Println("Server is running on port :8080")
// // 	http.ListenAndServe(":8080", nil)
// // }
