package api

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func Run() {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", GetTasks).Methods(http.MethodGet)
	router.HandleFunc("/tasks/", CreateTask).Methods(http.MethodPost)
	router.HandleFunc("/tasks/{id}", GetTaskByID).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{id}", UpdateTask).Methods(http.MethodPut)
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods(http.MethodDelete)

	router.HandleFunc("/categories", GetCategories).Methods(http.MethodGet)
	router.HandleFunc("/categories", CreateCategory).Methods(http.MethodPost)
	router.HandleFunc("/categories/{id}", DeleteCategory).Methods(http.MethodDelete)

	corsOptions := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	http.ListenAndServe(":8000", handlers.CORS(corsOptions, corsMethods, corsHeaders)(router))
}
