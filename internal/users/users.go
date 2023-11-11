package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        string    `db:"id"`
	LastName  string    `db:"last_name"`
	FirstName string    `db:"first_name"`
	BirthDate string    `db:"birth_date"`
	Gender    string    `db:"gender"`
	CreatedAt time.Time `db:"created_at"`
}

func HandleUsers(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	// HTTPメソッドで処理を振り分ける
	switch r.Method {
	case http.MethodGet:
		// GETリクエストの処理
		id := r.URL.Query().Get("id") // URLからidを取得
		var users []User
		var queryRows *sql.Rows

		if id != "" {
			// IDが存在するか確認
			var existingID string
			err := db.QueryRow("SELECT id FROM user_tbl WHERE id = $1", id).Scan(&existingID)
			if err == sql.ErrNoRows {
				http.Error(w, "Not Found: User ID not found", http.StatusNotFound)
				return
			} else if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// 特定のIDに一致するユーザー情報を取得
			rows, err := db.Query("SELECT * FROM user_tbl WHERE id = $1", id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer rows.Close()
			queryRows = rows
		} else {
			// すべてのユーザー情報を取得
			rows, err := db.Query("SELECT * FROM user_tbl")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer rows.Close()
			queryRows = rows
		}

		for queryRows.Next() {
			var user User
			err := queryRows.Scan(&user.Id, &user.LastName, &user.FirstName, &user.BirthDate, &user.Gender, &user.CreatedAt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		// POSTリクエストの処理
		var user User

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		birthdayTime, err := time.Parse("2006-01-02", user.BirthDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		birthday := birthdayTime.Format("2006-01-02T15:04:05Z")
		uuid := uuid.New()

		_, err = db.Exec("INSERT INTO user_tbl (id, last_name, first_name, birth_date, gender) VALUES ($1, $2, $3, $4, $5)", uuid, user.LastName, user.FirstName, birthday, user.Gender)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

	case http.MethodPut:
		// PUTリクエストの処理
		id := r.URL.Query().Get("id") // URLからidを取得
		if id == "" {
			http.Error(w, "Bad Request: Missing ID", http.StatusBadRequest)
			return
		}

		var user User

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// IDが存在するか確認
		var existingID string
		err := db.QueryRow("SELECT id FROM user_tbl WHERE id = $1", id).Scan(&existingID)
		if err == sql.ErrNoRows {
			http.Error(w, "Not Found: User ID not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		birthdayTime, err := time.Parse("2006-01-02", user.BirthDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		birthday := birthdayTime.Format("2006-01-02T15:04:05Z")

		_, err = db.Exec("UPDATE user_tbl SET last_name = $1, first_name = $2, birth_date = $3, gender = $4 WHERE id = $5", user.LastName, user.FirstName, birthday, user.Gender, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		// DELETEリクエストの処理
		id := r.URL.Query().Get("id") // URLからidを取得
		if id == "" {
			http.Error(w, "Bad Request: Missing ID", http.StatusBadRequest)
			return
		}

		// IDが存在するか確認
		var existingID string
		err := db.QueryRow("SELECT id FROM user_tbl WHERE id = $1", id).Scan(&existingID)
		if err == sql.ErrNoRows {
			http.Error(w, "Not Found: User ID not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("DELETE FROM user_tbl WHERE id = $1", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

}
