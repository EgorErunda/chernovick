package main

import (
	"fmt"
	//  "log"
	"net/http"
)

var (
	requestCount = 0
	fib1         = 0
	fib2         = 1
)

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	requestCount++
	fmt.Fprintf(w, "%d", fib1)
	fib1, fib2 = fib2, fib1+fib2
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", requestCount)
}

// func main() {
//  http.HandleFunc("/", FibonacciHandler)
//  http.HandleFunc("/metrics", MetricsHandler)
//  err := http.ListenAndServe(":8082", nil)
//  if err != nil {
//   log.Println("faile to start server", err)
//   return
//  }
//  log.Println("server was started")
// }
