package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sahilchauhan0603/backend/controllers"
	"github.com/sahilchauhan0603/backend/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitializeRoutes(router *mux.Router) {
    
	// Handle preflight requests for the /api/v1 endpoints
    router.PathPrefix("/api/v1").Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
        w.WriteHeader(http.StatusNoContent)
    })

	// Authentication routes
	// @Summary Microsoft login
	// @Description Initiate Microsoft login
	// @Tags Authentication
	// @Router /login [get]
	router.HandleFunc("/login", controllers.HandleMicrosoftLogin).Methods("GET")

	// @Summary Microsoft login callback
	// @Description Handle Microsoft login callback
	// @Tags Authentication
	// @Router /callback [get]
	router.HandleFunc("/callback", controllers.HandleMicrosoftCallback).Methods("GET")

	// Swagger UI route
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)


	// Uploader routes
	uploaderRouter := router.PathPrefix("/api/v1").Subrouter()
	uploaderRouter.Use(middleware.JWTVerify)

	// @Summary Create a new uploader
	// @Description Create a new uploader
	// @Tags Uploaders
	// @Accept json
	// @Produce json
	// @Param uploader body models.Uploader true "Uploader information"
	// @Success 201 {object} models.Uploader
	// @Failure 400 {string} string "Bad Request"
	// @Router /api/v1/upload [post]
	uploaderRouter.HandleFunc("/upload", controllers.CreateUploader).Methods("POST")

	// @Summary Get all uploaders
	// @Description Retrieve all uploaders
	// @Tags Uploaders
	// @Produce json
	// @Success 200 {array} models.Uploader
	// @Failure 500 {string} string "Internal Server Error"
	// @Router /api/v1/upload [get]
	uploaderRouter.HandleFunc("/upload", controllers.GetUploader).Methods("GET")

	// @Summary Get uploader by ID
	// @Description Retrieve uploader by ID
	// @Tags Uploaders
	// @Produce json
	// @Param id path int true "Uploader ID"
	// @Success 200 {object} models.Uploader
	// @Failure 404 {string} string "Uploader not found"
	// @Router /api/v1/upload/{id} [get]
	uploaderRouter.HandleFunc("/upload/{id}", controllers.GetUploaderID).Methods("GET")

	// @Summary Update uploader
	// @Description Update uploader by ID
	// @Tags Uploaders
	// @Accept json
	// @Produce json
	// @Param id path int true "Uploader ID"
	// @Param uploader body models.Uploader true "Uploader information"
	// @Success 200 {object} models.Uploader
	// @Failure 404 {string} string "Uploader not found"
	// @Failure 400 {string} string "Bad Request"
	// @Router /api/v1/upload/{id} [put]
	uploaderRouter.HandleFunc("/upload/{id}", controllers.UpdateUploader).Methods("PUT")

	// @Summary Delete uploader
	// @Description Delete uploader by ID
	// @Tags Uploaders
	// @Param id path int true "Uploader ID"
	// @Success 204 {string} string "No Content"
	// @Failure 404 {string} string "Uploader not found"
	// @Router /api/v1/upload/{id} [delete]
	uploaderRouter.HandleFunc("/upload/{id}", controllers.DeleteUploader).Methods("DELETE")


	// Work routes
	workRouter := router.PathPrefix("/api/v1").Subrouter()
	workRouter.Use(middleware.JWTVerify)

	// @Summary Create a new work
	// @Description Create a new work
	// @Tags Works
	// @Accept json
	// @Produce json
	// @Param project body models.Work true "Work information"
	// @Success 201 {object} models.Work
	// @Failure 400 {string} string "Bad Request"
	// @Router /api/v1/work [post]
	workRouter.HandleFunc("/work", controllers.CreateWork).Methods("POST")

	// @Summary Get all works
	// @Description Retrieve all works
	// @Tags Works
	// @Produce json
	// @Success 200 {array} models.Work
	// @Failure 500 {string} string "Internal Server Error"
	// @Router /api/v1/work [get]
	workRouter.HandleFunc("/work", controllers.GetWork).Methods("GET")

	// @Summary Get work by ID
	// @Description Retrieve work by ID
	// @Tags Works
	// @Produce json
	// @Param id path int true "Work ID"
	// @Success 200 {object} models.Work
	// @Failure 404 {string} string "Work not found"
	// @Router /api/v1/work/{id} [get]
	workRouter.HandleFunc("/work/{id}", controllers.GetWorkID).Methods("GET")

	// @Summary Update work
	// @Description Update work by ID
	// @Tags Works
	// @Accept json
	// @Produce json
	// @Param id path int true "Work ID"
	// @Param project body models.Work true "Work information"
	// @Success 200 {object} models.Work
	// @Failure 404 {string} string "Work not found"
	// @Failure 400 {string} string "Bad Request"
	// @Router /api/v1/work/{id} [put]
	workRouter.HandleFunc("/work/{id}", controllers.UpdateWork).Methods("PUT")

	// @Summary Delete work
	// @Description Delete work by ID
	// @Tags Works
	// @Param id path int true "Work ID"
	// @Success 204 {string} string "No Content"
	// @Failure 404 {string} string "Work not found"
	// @Router /api/v1/work/{id} [delete]
	workRouter.HandleFunc("/work/{id}", controllers.DeleteWork).Methods("DELETE")


	// Admin routes
	adminRouter := router.PathPrefix("/api/v1").Subrouter()
	adminRouter.Use(middleware.JWTVerify)

	// @Summary Create a new admin
	// @Description Create a new admin
	// @Tags Admins
	// @Accept json
	// @Produce json
	// @Param admin body models.Admin true "Admin information"
	// @Success 201 {object} models.Admin
	// @Failure 400 {string} string "Bad Request"
	// @Router /api/v1/admin [post]
	adminRouter.HandleFunc("/admin", controllers.CreateAdmin).Methods("POST")

	// @Summary Get all admins
	// @Description Retrieve all admins
	// @Tags Admins
	// @Produce json
	// @Success 200 {array} models.Admin
	// @Failure 500 {string} string "Internal Server Error"
	// @Router /api/v1/admin [get]
	adminRouter.HandleFunc("/admin", controllers.GetAdmin).Methods("GET")

	// @Summary Get admin by ID
	// @Description Retrieve admin by ID
	// @Tags Admins
	// @Produce json
	// @Param id path int true "Admin ID"
	// @Success 200 {object} models.Admin
	// @Failure 404 {string} string "Admin not found"
	// @Router /api/v1/admin/{id} [get]
	adminRouter.HandleFunc("/admin/{id}", controllers.GetAdminID).Methods("GET")

	// @Summary Update admin
	// @Description Update admin by ID
	// @Tags Admins
	// @Accept json
	// @Produce json
	// @Param id path int true "Admin ID"
	// @Param admin body models.Admin true "Admin information"
	// @Success 200 {object} models.Admin
	// @Failure 404 {string} string "Admin not found"
	// @Failure 400 {string} string "Bad Request"
	// @Router /api/v1/admin/{id} [put]
	adminRouter.HandleFunc("/admin/{id}", controllers.UpdateAdmin).Methods("PUT")

	// @Summary Delete admin
	// @Description Delete admin by ID
	// @Tags Admins
	// @Param id path int true "Admin ID"
	// @Success 204 {string} string "No Content"
	// @Failure 404 {string} string "Admin not found"
	// @Router /api/v1/admin/{id} [delete]
	adminRouter.HandleFunc("/admin/{id}", controllers.DeleteAdmin).Methods("DELETE")


	// testing route
	emailRouter := router.PathPrefix("/api/v1").Subrouter()
	emailRouter.Use(middleware.JWTVerify)

	// @Summary Create a email
	// @Description Create a email
	// @Tags Email
	// @Accept json
	// @Produce json
	// @Param admin body models.EmailRequest true "Email sending"
	// @Success 201 {object} models.EmailRequest
	// @Failure 400 {string} string "Bad Request"
	// @Router /api/v1/sendEmail [post]
    emailRouter.HandleFunc("/sendEmail", controllers.SendEmail).Methods("POST")

}
