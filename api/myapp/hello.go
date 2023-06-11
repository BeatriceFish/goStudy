package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
}

// Handler 处理根路径的 GET 请求
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	message := "Hello, world!"
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "%s"}`, message)
}
