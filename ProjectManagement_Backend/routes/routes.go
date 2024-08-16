package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sahilchauhan0603/backend/controllers"
)

func InitializeRoutes(router *mux.Router) {

	// Handle preflight requests for the /api/v1 endpoints
	router.PathPrefix("/api/v1").Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.Header().Set("Access-control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusNoContent)
	})

	router.HandleFunc("/login", controllers.HandleMicrosoftLogin).Methods("GET")
	router.HandleFunc("/callback", controllers.HandleMicrosoftCallback).Methods("GET")

	// Uploader routes
	r := router.PathPrefix("/api/v1").Subrouter()
	// r.Use(middleware.JWTVerify)

	r.HandleFunc("/upload", controllers.CreateUploader).Methods("POST")
	r.HandleFunc("/upload", controllers.GetUploader).Methods("GET")
	r.HandleFunc("/upload/{enrollmentNo}", controllers.GetUploaderID).Methods("GET")
	r.HandleFunc("/upload/{enrollmentNo}", controllers.UpdateUploader).Methods("PUT")
	r.HandleFunc("/upload/{enrollmentNo}", controllers.DeleteUploader).Methods("DELETE")

	//work routes
	r.HandleFunc("/work", controllers.CreateWork).Methods("POST")
	r.HandleFunc("/work", controllers.GetWork).Methods("GET")
	r.HandleFunc("/work/{workID}", controllers.GetWorkID).Methods("GET")
	r.HandleFunc("/work/{workID}", controllers.UpdateWork).Methods("PUT")
	r.HandleFunc("/work/{workID}", controllers.DeleteWork).Methods("DELETE")

	// Admin routes
	r.HandleFunc("/admin", controllers.CreateAdmin).Methods("POST")
	r.HandleFunc("/admin", controllers.GetAdmin).Methods("GET")
	r.HandleFunc("/admin/{adminID}", controllers.GetAdminID).Methods("GET")
	r.HandleFunc("/admin/{adminID}", controllers.UpdateAdmin).Methods("PUT")
	r.HandleFunc("/admin/{adminID}", controllers.DeleteAdmin).Methods("DELETE")

	//Email route
	r.HandleFunc("/contact", controllers.ContactUSHandler).Methods("POST")

}
