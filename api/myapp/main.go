package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", helloHandler)

	log.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

// helloHandler 处理 /api/hello 路径的 GET 请求
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	message := "Hello, world!"
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "%s"}`, message)
}
