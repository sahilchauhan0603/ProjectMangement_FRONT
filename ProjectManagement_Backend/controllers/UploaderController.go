package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	database "github.com/sahilchauhan0603/backend/config"
	models "github.com/sahilchauhan0603/backend/models"
)

func CreateUploader(w http.ResponseWriter, r *http.Request) {
	var uploader models.ProjectUploader
	if err := json.NewDecoder(r.Body).Decode(&uploader); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if result := database.DB.Create(&uploader); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uploader)
}

func GetUploader(w http.ResponseWriter, r *http.Request) {
	var uploaders []models.ProjectUploader
	if result := database.DB.Find(&uploaders); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uploaders)
}

func GetUploaderID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	enrollmentNo, err := strconv.Atoi(params["enrollmentNo"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var uploader []models.ProjectUploader
	if result := database.DB.First(&uploader, enrollmentNo); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uploader)
}


func UpdateUploader(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	enrollmentNo, err := strconv.Atoi(params["enrollmentNo"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var uploader models.ProjectUploader
	if result := database.DB.First(&uploader, enrollmentNo); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&uploader); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&uploader)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(uploader)
}

func DeleteUploader(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	enrollmentNo, err := strconv.Atoi(params["enrollmentNo"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if result := database.DB.Delete(&models.ProjectUploader{}, enrollmentNo); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Role successfully deleted"})

}
