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

	return r
}
