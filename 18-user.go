package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	ID    int
	Name  string
	Email string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", listUsersHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	http.ListenAndServe(":8080", mux)
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []*user{}
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, &u)
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var u user
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Erros decoding JSON", http.StatusBadRequest)
		return
	}

	if _, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", u.Name, u.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
