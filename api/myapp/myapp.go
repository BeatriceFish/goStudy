package handler

import (
	"fmt"
	"net/http"
)

// Handler 是一个适用于 Vercel Serverless 函数的 HTTP 处理函数。
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from your Go app!")
}
