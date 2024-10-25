package gormDatabase

import (
	"net/http"

	"github.com/gorilla/mux"
)

var db = Connect()

func Run() {
	//CreateUser("Aknur", 20, "Hello!", "profile_image.png")
	//GetUsers()
	//UpdateUser(3, "Aknur", 21)
	//DeleteUser(2)
	//UpdateProfile(2, "Updated Bio for Ainur", "new_photo2.jpg")

	router := mux.NewRouter()
	router.Use(corsMiddleware)

	router.HandleFunc("/users", GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", DeleteUser).Methods(http.MethodDelete)

	http.ListenAndServe(":8080", router)

}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}
