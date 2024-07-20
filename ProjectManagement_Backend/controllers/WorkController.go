package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/backend/config"
	models "github.com/sahilchauhan0603/backend/models"
)

// CreateWork creates a new Work
// @Summary Create a new Work
// @Description Create a new work
// @Tags work
// @Accept json
// @Produce json
// @Param work body models.Work true "Work"
// @Success 201 {object} models.Work
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/work [post]
func CreateWork(w http.ResponseWriter, r *http.Request) {
	var work models.Work
	if err := json.NewDecoder(r.Body).Decode(&work); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.Work{}) {
		if err := database.DB.AutoMigrate(&models.Work{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if result := database.DB.Create(&work); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(work)
}

// GetWork retrieves all works
// @Summary Get all works
// @Description Get all works
// @Tags work
// @Produce json
// @Success 200 {array} models.Work
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/work [get]
func GetWork(w http.ResponseWriter, r *http.Request) {
	var work []models.Work
	if result := database.DB.Find(&work); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(work)
}

// GetWorkID retrieves a work by ID
// @Summary Get work by ID
// @Description Get work by ID
// @Tags work
// @Produce json
// @Param id path int true "Work ID"
// @Success 200 {object} models.Work
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Not Found"
// @Router /api/v1/work/{id} [get]
func GetWorkID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var work models.Work
	if result := database.DB.First(&work, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(work)
}

// UpdateWork updates a work by ID
// @Summary Update work by ID
// @Description Update work by ID
// @Tags work
// @Accept json
// @Produce json
// @Param id path int true "Work ID"
// @Param work body models.Work true "Work"
// @Success 200 {object} models.Work
// @Failure 400 {string} string "Invalid ID or Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /api/v1/work/{id} [put]
func UpdateWork(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var work models.Work
	if result := database.DB.First(&work, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&work); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&work)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(work)
}

// DeleteWork deletes a work by ID
// @Summary Delete work by ID
// @Description Delete work by ID
// @Tags work
// @Param id path int true "Work ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Invalid ID"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/work/{id} [delete]
func DeleteWork(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if result := database.DB.Delete(&models.Work{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
