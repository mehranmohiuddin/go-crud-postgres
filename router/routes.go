package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mehranmohiuddin/go-crud-postgres/models"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	}).Methods("GET")

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error oading .env file")
		}

		db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
		if err != nil {
			log.Fatal("failed to open a db connection:", err)
		}
		defer db.Close()

		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			log.Fatal("Failed to execute query: ", err)
		}
		defer rows.Close()

		var users []models.User
		for rows.Next() {
			var user models.User

			err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location)
			if err != nil {
				log.Fatalf("Unable to scan the row. %v", err)
			}

			users = append(users, user)
		}

		json.NewEncoder(w).Encode(users)
	}).Methods("GET")

	return r
}
