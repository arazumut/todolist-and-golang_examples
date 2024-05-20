package main

import (
	"log"
	"net/http"
	"time"

	"example.com/todo-app-go/internal/api"
	"example.com/todo-app-go/internal/auth"
	"example.com/todo-app-go/internal/mock"
)

func main() {
	
	mockService := mock.NewMockService()

	
	authService := auth.NewAuthService()

	
	apiHandler := api.NewAPIHandler(mockService, authService)

	// Yönlendirilecek rotaları tanımla
	http.HandleFunc("/login", apiHandler.Login)
	http.HandleFunc("/todo/create", apiHandler.CreateToDoList)
	http.HandleFunc("/todo/delete", apiHandler.DeleteToDoList)
	http.HandleFunc("/todo/addmessage", apiHandler.AddToDoMessage)
	http.HandleFunc("/todo/deletemessage", apiHandler.DeleteToDoMessage)
	http.HandleFunc("/todo/updatemessage", apiHandler.UpdateToDoMessage)
	http.HandleFunc("/todo/get", apiHandler.GetToDoLists)

	// HTTP sunucusunu başlat
	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil, // DefaultMux kullan
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server listening on port 8080...")
	log.Fatal(server.ListenAndServe())
}
