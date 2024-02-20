package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// response := Response{
	// 	Status: "Success",
	// 	Code:   200,
	// }

	// 设置响应头的内容类型为 JSON
	w.Header().Set("Content-Type", "application/json")

	// 将 response 结构体编码为 JSON 并写入响应
	json.NewEncoder(w).Encode(Response{
		Status: "Success",
		Code:   200,
	})
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
