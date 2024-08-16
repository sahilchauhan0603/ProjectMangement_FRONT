package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/backend/config"
	models "github.com/sahilchauhan0603/backend/models"
)

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var admin models.ProjectAdmin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.ProjectAdmin{}) {
		if err := database.DB.AutoMigrate(&models.ProjectAdmin{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if result := database.DB.Create(&admin); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(admin)
}

func GetAdmin(w http.ResponseWriter, r *http.Request) {
	var admin []models.ProjectAdmin
	if result := database.DB.Find(&admin); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(admin)
}

func GetAdminID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	adminID, err := strconv.Atoi(params["adminID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var admin []models.ProjectAdmin
	if result := database.DB.First(&admin, adminID); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(admin)
}

func UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	adminID, err := strconv.Atoi(params["adminID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var admin models.ProjectAdmin
	if result := database.DB.First(&admin, adminID); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&admin)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(admin)
}

func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	adminID, err := strconv.Atoi(params["adminID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if result := database.DB.Delete(&models.ProjectAdmin{}, adminID); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Role successfully deleted"})
}
