package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	database "github.com/sahilchauhan0603/backend/config"
	routes "github.com/sahilchauhan0603/backend/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database connection
	database.DatabaseConnector()

	// Create a new router
	router := mux.NewRouter()
	routes.InitializeRoutes(router)

	// Enable CORS
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Change to specific origins in production
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
		handlers.AllowCredentials(),
	)

	// Set the port for the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	// Print server start message
	fmt.Printf("Server is running on port %s\n", port)

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), cors(router)))
}
