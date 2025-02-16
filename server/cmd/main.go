package main

import (
	"backend/blog/configs"
	"backend/blog/internal/auth"
	"fmt"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	router := http.NewServeMux()

	//Repository

	//Services
	authService := auth.NewAuthService()

	//Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		AuthService: authService,
		Configs:     config,
	})

	server := http.Server{
		Addr:    "8081",
		Handler: router,
	}

	fmt.Println("Сервер запущен на порту 8081")
	server.ListenAndServe()
}
