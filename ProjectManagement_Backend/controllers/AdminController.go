package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/backend/config"
	models "github.com/sahilchauhan0603/backend/models"
)

// CreateAdmin creates a new admin
// @Summary Create a new admin
// @Description Create a new admin
// @Tags admin
// @Accept json
// @Produce json
// @Param admin body models.Admin true "Admin"
// @Success 201 {object} models.Admin
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/admin [post]
func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var admin models.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.Admin{}) {
		if err := database.DB.AutoMigrate(&models.Admin{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if result := database.DB.Create(&admin); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(admin)
}

// GetAdmin retrieves all admins
// @Summary Get all admins
// @Description Get all admins
// @Tags admin
// @Produce json
// @Success 200 {array} models.Admin
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/admin [get]
func GetAdmin(w http.ResponseWriter, r *http.Request) {
	var admin []models.Admin
	if result := database.DB.Find(&admin); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(admin)
}

// GetAdminID retrieves an admin by ID
// @Summary Get admin by ID
// @Description Get admin by ID
// @Tags admin
// @Produce json
// @Param id path int true "Admin ID"
// @Success 200 {object} models.Admin
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Not Found"
// @Router /api/v1/admin/{id} [get]
func GetAdminID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var admin models.Admin
	if result := database.DB.First(&admin, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(admin)
}

// UpdateAdmin updates an admin by ID
// @Summary Update admin by ID
// @Description Update admin by ID
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Admin ID"
// @Param admin body models.Admin true "Admin"
// @Success 200 {object} models.Admin
// @Failure 400 {string} string "Invalid ID or Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /api/v1/admin/{id} [put]
func UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var admin models.Admin
	if result := database.DB.First(&admin, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&admin)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(admin)
}

// DeleteAdmin deletes an admin by ID
// @Summary Delete admin by ID
// @Description Delete admin by ID
// @Tags admin
// @Param id path int true "Admin ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Invalid ID"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/admin/{id} [delete]
func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if result := database.DB.Delete(&models.Admin{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
