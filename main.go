package main

import (
	"fmt"
	"net/http"

	"github.com/mehranmohiuddin/go-crud-postgres/router"
)

func main() {
	fmt.Println("Running on port 8080")

	r := router.Router()

	http.ListenAndServe(":8080", r)
}
