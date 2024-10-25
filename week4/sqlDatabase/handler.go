package sqlDatabase

import (
	"github.com/gorilla/mux"
	"net/http"
)

var db = Connect()

func Run() {
	router := mux.NewRouter()
	router.Use(corsMiddleware)

	router.HandleFunc("/users", GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", DeleteUser).Methods(http.MethodDelete)

	//users := []UserForm{
	//	{Name: "A", Age: 21},
	//	{Name: "B", Age: 20},
	//	{Name: "C", Age: 22},
	//}
	//InsertMultipleUsers(users)

	http.ListenAndServe(":8080", router)
}
