package main

import (
	"fmt"
	"go/adv-demo/internal/auth"
	"net/http"
)

func main() {
	// _ := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router)
	// auth/login
	// auth/register

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
