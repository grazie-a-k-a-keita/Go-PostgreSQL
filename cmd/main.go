package main

import (
	"go-practice/internal/users"
	"go-practice/internal/utility"
	"log"
	"net/http"
)

func main() {

	db, err := utility.ConnectionDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	// ~:8080/users
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users.HandleUsers(w, r, db)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
