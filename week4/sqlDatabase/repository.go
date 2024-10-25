package sqlDatabase

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	UserID  uint   `json:"user_id"`
}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   i

func jsonRespond(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonRespond(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}
	query := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`
	var userId uint
	err := db.QueryRow(query, user.Name, user.Age).Scan(&userId)
	if err != nil {
		jsonRespond(w, http.StatusInternalServerError,
			ErrorResponse{Error: fmt.Sprintf("Failed to create user: %v", err)})
		return
	}
	jsonRespond(w, http.StatusCreated, SuccessResponse{
		Message: "User created successfully",
		UserID:  userId,
	})
}

func CreateUser3(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonRespond(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}
	var existingUserID uint
	checkQuery := `SELECT id FROM users WHERE name = $1`
	err := db.QueryRow(checkQuery, user.Name).Scan(&existingUserID)
	if err == nil {
		jsonRespond(w, http.StatusConflict, ErrorResponse{Error: "A user with this name already exists"})
		return
	}
	query := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`
	var userId uint
	err = db.QueryRow(query, user.Name, user.Age).Scan(&userId)
	if err != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: fmt.Sprintf("Failed to create user: %v", err)})
		return
	}

	query = `INSERT INTO profiles (user_id, bio, profile_picture_url) VALUES ($1, $2, $3)`
	_, err = db.Exec(query, userId, "", "")
	if err != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: fmt.Sprintf("Failed to create profile: %v", err)})
		return
	}
	jsonRespond(w, http.StatusCreated, SuccessResponse{
		Message: "User created successfully",
		UserID:  userId,
	})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	minAge := r.URL.Query().Get("min_age")
	maxAge := r.URL.Query().Get("max_age")

	pageNum := 1
	limitNum := 5

	if p, err := strconv.Atoi(page); err == nil && p > 0 {
		pageNum = p
	}
	if l, err := strconv.Atoi(limit); err == nil && l > 0 {
		limitNum = l
	}
	offset := (pageNum - 1) * limitNum

	countQuery := `SELECT COUNT(*) FROM users WHERE 1=1`
	var total int
	err := db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		http.Error(w, "Failed to retrieve user count", http.StatusInternalServerError)
		return
	}

	query := `SELECT id, name, age FROM users WHERE 1=1`
	if minAge != "" {
		query += ` AND age >= ` + minAge
	}
	if maxAge != "" {
		query += ` AND age <= ` + maxAge
	}
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limitNum, offset)

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			continue
		}
		users = append(users, user)
	}
	response := map[string]interface{}{
		"data":  users,
		"total": total,
		"page":  pageNum,
	}
	jsonRespond(w, http.StatusOK, response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	id := getId(w, r)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonRespond(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}
	query := `UPDATE users SET name = $1, age = $2 WHERE id = $3`
	_, err := db.Exec(query, user.Name, user.Age, id)
	if err != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: fmt.Sprintf("Failed to update user: %v", err)})
		return
	}
	jsonRespond(w, http.StatusOK, SuccessResponse{
		Message: "User updated successfully",
		UserID:  id,
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := getId(w, r)

	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: fmt.Sprintf("Failed to delete user: %v", err)})
		return
	}
	jsonRespond(w, http.StatusOK, SuccessResponse{
		Message: "User deleted successfully",
		UserID:  id,
	})
}

func InsertMultipleUsers(users []UserForm) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("Failed to begin transaction: %v\n", err)
		return
	}
	userQuery := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`

	for _, user := range users {
		var userID int
		err = tx.QueryRow(userQuery, user.Name, user.Age).Scan(&userID)
		if err != nil {
			tx.Rollback()
			fmt.Printf("Failed to add user %s: %v\n", user.Name, err)
			return
		}
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		fmt.Printf("Failed to commit transaction: %v\n", err)
		return
	}
	fmt.Println("Users inserted successfully")
}

func getId(w http.ResponseWriter, r *http.Request) uint {
	vars := mux.Vars(r)
	idStr := vars["id"]
	idUint, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		jsonRespond(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid ID. Must be a number"})
		return 0
	}
	return uint(idUint)
}
