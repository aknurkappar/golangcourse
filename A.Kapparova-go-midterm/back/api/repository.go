package api

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
	ID      uint   `json:"id"`
}

func jsonRespond(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title      string `json:"title"`
		CategoryID *uint  `json:"category_id,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Title == "" {
		jsonRespond(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid input or missing title"})
		return
	}
	task := Task{
		Title:      input.Title,
		Status:     "todo",
		CategoryID: input.CategoryID,
	}
	if err := DB.Create(&task).Error; err != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: fmt.Sprintf("Failed to create task: %v", err)})
		return
	}
	jsonRespond(w, http.StatusCreated, SuccessResponse{Message: "Task created successfully", ID: task.ID})
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	taskIDParam := mux.Vars(r)["id"]
	taskID, err := strconv.ParseUint(taskIDParam, 10, 32)
	if err != nil {
		jsonRespond(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid task ID"})
		return
	}

	var task Task
	if err := DB.First(&task, taskID).Error; err != nil {
		jsonRespond(w, http.StatusNotFound, ErrorResponse{Error: "Task not found"})
		return
	}

	jsonRespond(w, http.StatusOK, task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	categoryIDParam := r.URL.Query().Get("category_id")

	var tasks []Task
	query := DB.Model(&Task{}).Preload("Category")
	if categoryIDParam != "" {
		categoryID, err := strconv.ParseUint(categoryIDParam, 10, 32)
		if err != nil {
			jsonRespond(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid category ID"})
			return
		}
		query = query.Where("category_id = ?", categoryID)
	}
	if err := query.Find(&tasks).Error; err != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: "Failed to retrieve tasks"})
		return
	}
	jsonRespond(w, http.StatusOK, tasks)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskIDParam := mux.Vars(r)["id"]
	taskID, err := strconv.ParseUint(taskIDParam, 10, 32)
	if err != nil {
		jsonRespond(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid task ID"})
		return
	}
	var input struct {
		Title      string `json:"title"`
		Status     string `json:"status"`
		CategoryID *uint  `json:"category_id,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		jsonRespond(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}
	defer r.Body.Close()

	var task Task
	if err := DB.First(&task, taskID).Error; err != nil {
		jsonRespond(w, http.StatusNotFound, ErrorResponse{Error: "Task not found"})
		return
	}

	task.Title = input.Title
	task.Status = input.Status
	task.CategoryID = input.CategoryID

	if err := DB.Save(&task).Error; err != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: "Error updating task: " + err.Error()})
		return
	}

	jsonRespond(w, http.StatusOK, SuccessResponse{Message: "Task updated successfully", ID: task.ID})
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		jsonRespond(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid task ID"})
		return
	}
	if err := DB.Delete(&Task{}, taskID).Error; err != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: "Error deleting task: " + err.Error()})
		return
	}
	jsonRespond(w, http.StatusOK, SuccessResponse{Message: "Task deleted successfully", ID: uint(taskID)})
}

// Category

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		jsonRespond(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}

	category := Category{Name: input.Name}
	if err := DB.Create(&category).Error; err != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating category: " + err.Error()})
		return
	}

	jsonRespond(w, http.StatusCreated, SuccessResponse{Message: "Category created successfully", ID: category.ID})
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []Category
	if err := DB.Find(&categories).Error; err != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: "Error retrieving categories: " + err.Error()})
		return
	}

	jsonRespond(w, http.StatusOK, categories)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		jsonRespond(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid category ID"})
		return
	}

	if err := DB.Delete(&Category{}, categoryID).Error; err != nil {
		jsonRespond(w, http.StatusInternalServerError, ErrorResponse{Error: "Error deleting category: " + err.Error()})
		return
	}

	jsonRespond(w, http.StatusOK, SuccessResponse{Message: "Category deleted successfully", ID: uint(categoryID)})
}
