package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	server := http.Server{
		Addr:    "8081",
		Handler: router,
	}

	fmt.Println("Сервер запущен на порту 8081")
	server.ListenAndServe()
}
