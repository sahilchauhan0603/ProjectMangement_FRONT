package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/backend/config"
	models "github.com/sahilchauhan0603/backend/models"
)

func CreateWork(w http.ResponseWriter, r *http.Request) {
	var work models.ProjectWork
	if err := json.NewDecoder(r.Body).Decode(&work); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.ProjectWork{}) {
		if err := database.DB.AutoMigrate(&models.ProjectWork{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if result := database.DB.Create(&work); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(work)
}

func GetWork(w http.ResponseWriter, r *http.Request) {
	var work []models.ProjectWork
	if result := database.DB.Find(&work); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(work)
}

func GetWorkID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workID, err := strconv.Atoi(params["workID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var work []models.ProjectWork
	if result := database.DB.First(&work, workID); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(work)
}

func UpdateWork(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workID, err := strconv.Atoi(params["workID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var work models.ProjectWork
	if result := database.DB.First(&work, workID); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&work); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&work)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(work)
}

func DeleteWork(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workID, err := strconv.Atoi(params["workID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if result := database.DB.Delete(&models.ProjectWork{}, workID); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Role successfully deleted"})
}
