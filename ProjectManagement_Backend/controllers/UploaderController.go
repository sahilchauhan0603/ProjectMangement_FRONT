package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	database "github.com/sahilchauhan0603/backend/config"
	models "github.com/sahilchauhan0603/backend/models"
)

// CreateUploader creates a new uploader
// @Summary Create a new uploader
// @Description Create a new uploader
// @Tags uploader
// @Accept json
// @Produce json
// @Param uploader body models.Uploader true "Uploader"
// @Success 201 {object} models.Uploader
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/upload [post]
func CreateUploader(w http.ResponseWriter, r *http.Request) {
	var uploader models.Uploader
	if err := json.NewDecoder(r.Body).Decode(&uploader); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if result := database.DB.Create(&uploader); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(uploader)
}

// GetUploader retrieves all uploaders
// @Summary Get all uploaders
// @Description Get all uploaders
// @Tags uploader
// @Produce json
// @Success 200 {array} models.Uploader
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/upload [get]
func GetUploader(w http.ResponseWriter, r *http.Request) {
	var uploaders []models.Uploader
	if result := database.DB.Find(&uploaders); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(uploaders)
}

// GetUploaderID retrieves an uploader by ID
// @Summary Get uploader by ID
// @Description Get uploader by ID
// @Tags uploader
// @Produce json
// @Param id path int true "Uploader ID"
// @Success 200 {object} models.Uploader
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Not Found"
// @Router /api/v1/upload/{id} [get]
func GetUploaderID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var uploader models.Uploader
	if result := database.DB.First(&uploader, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(uploader)
}

// UpdateUploader updates an uploader by ID
// @Summary Update uploader by ID
// @Description Update uploader by ID
// @Tags uploader
// @Accept json
// @Produce json
// @Param id path int true "Uploader ID"
// @Param uploader body models.Uploader true "Uploader"
// @Success 200 {object} models.Uploader
// @Failure 400 {string} string "Invalid ID or Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /api/v1/upload/{id} [put]
func UpdateUploader(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var uploader models.Uploader
	if result := database.DB.First(&uploader, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&uploader); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&uploader)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(uploader)
}

// DeleteUploader deletes an uploader by ID
// @Summary Delete uploader by ID
// @Description Delete uploader by ID
// @Tags uploader
// @Param id path int true "Uploader ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Invalid ID"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/upload/{id} [delete]
func DeleteUploader(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if result := database.DB.Delete(&models.Uploader{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
