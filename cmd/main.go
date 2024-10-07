package main

import (
	"go-practice/internal/users"
	"go-practice/internal/utility"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func main() {

	db, err := utility.ConnectionDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ~:8080/users
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users.HandleUsers(w, r, db, "")
	})

	// ~:8080/users/{id}
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		sub := strings.TrimPrefix(r.URL.Path, "/products")
		_, id := filepath.Split(sub)
		if id != "" {
			users.HandleUsers(w, r, db, id)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
