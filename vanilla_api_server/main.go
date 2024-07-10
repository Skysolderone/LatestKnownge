package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /hello", Hello)
	authRouter := http.NewServeMux()
	// authRouter.HandleFunc("PUT /hello/world", Addworld)
	router.Handle("/", CheckAuth(authRouter))
	server := http.Server{
		Addr:    ":8080",
		Handler: Logging(router),
	}
	server.ListenAndServe()
}

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("token")
		isValid := headerToken == "secretTOKEN"
		if !isValid {
			fmt.Println("invalid token:", headerToken)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Hello(w http.ResponseWriter, r *http.Request) {
	message := "Hi there"
	w.Write([]byte(message))
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		serveDuaration := time.Since(startTime)
		fmt.Println(r.Method, r.URL.Path, serveDuaration)
	})
}
