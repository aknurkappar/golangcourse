package gormDatabase

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	UserID  uint   `json:"user_id"`
}

type UserForm struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func jsonRespond(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
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

func pagination(r *http.Request) (int, int, int) {
	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(q.Get("limit"))
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	return offset, limit, page
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var input UserForm
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		jsonRespond(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}
	var existingUser User
	if err := db.Where("name = ?", input.Name).First(&existingUser).Error; err == nil {
		jsonRespond(w, http.StatusConflict, ErrorResponse{Error: "User with this name already exists"})
		return
	}
	user := User{
		Name: input.Name,
		Age:  input.Age,
		Profile: Profile{
			Bio:               " ",
			ProfilePictureURL: " ",
		},
	}
	if err := db.Create(&user).Error; err != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: "Failed to create user and profile"})
	}
	jsonRespond(w, http.StatusCreated, SuccessResponse{Message: "User created", UserID: user.ID})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := getId(w, r)
	if userID == 0 {
		return
	}
	var input UserForm

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		jsonRespond(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		jsonRespond(w, http.StatusNotFound, ErrorResponse{Error: "User not found"})
		return
	}
	var existingUser User
	if err := db.Where("name = ? AND id != ?", input.Name, userID).First(&existingUser).Error; err == nil {
		jsonRespond(w, http.StatusConflict, ErrorResponse{Error: "User with this name already exists"})
		return
	}
	if err := db.Model(&user).Updates(User{Name: input.Name, Age: input.Age}).Error; err != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: "Failed to update user"})
		return
	}
	jsonRespond(w, http.StatusOK, SuccessResponse{Message: "User updated", UserID: userID})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	offset, limit, page := pagination(r)
	result := db.Preload("Profile").Offset(offset).Limit(limit).Find(&users)

	if result.Error != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: "Failed to get users"})
		return
	}

	totalUsers := int64(0)
	db.Model(&User{}).Count(&totalUsers)

	response := struct {
		Users      []User `json:"users"`
		TotalCount int64  `json:"total_count"`
		Page       int    `json:"page"`
		Limit      int    `json:"limit"`
	}{
		Users:      users,
		TotalCount: totalUsers,
		Page:       page,
		Limit:      limit,
	}

	jsonRespond(w, http.StatusOK, response)
}

//func GetUsers(w http.ResponseWriter, r *http.Request) {
//	var users []User
//	if err := db.Preload("Profile").Find(&users).Error; err != nil {
//		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: "Failed to get users"})
//		return
//	}
//	jsonRespond(w, http.StatusOK, users)
//}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID := getId(w, r)
	if userID == 0 {
		return
	}
	var user struct {
		Bio               string `json:"bio"`
		ProfilePictureURL string `json:"profile_picture_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonRespond(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}
	var profile Profile
	if err := db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		jsonRespond(w, http.StatusNotFound, ErrorResponse{Error: "Profile not found"})
		return
	}
	profile.Bio = user.Bio
	profile.ProfilePictureURL = user.ProfilePictureURL
	if err := db.Save(&profile).Error; err != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: "Failed to update profile"})
		return
	}
	jsonRespond(w, http.StatusOK, SuccessResponse{Message: "Profile updated", UserID: userID})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := getId(w, r)
	if userID == 0 {
		return
	}
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", userID).Delete(&Profile{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&User{}, userID).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: "Failed to delete user and profile"})
		return
	}
	jsonRespond(w, http.StatusOK, SuccessResponse{Message: "User deleted", UserID: userID})
}
