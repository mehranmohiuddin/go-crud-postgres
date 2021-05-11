package router

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mehranmohiuddin/go-crud-postgres/handlers"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.BaseHandler).Methods("GET")
	r.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")

	return r
}
