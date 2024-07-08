package main

import (
	"log"
	"net/http"
	"todo-api/database"
	"todo-api/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database connection using environment variables
	database.Init()

	defer database.Close()

	router := mux.NewRouter()

	router.HandleFunc("/todo", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/todo", handlers.GetTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", handlers.GetTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", handlers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{id}", handlers.DeleteTodo).Methods("DELETE")

	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
