package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	models "github.com/sahilchauhan0603/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnector() {
	// Retrieve environment variables for the database
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbName == "" || dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" {
		log.Fatal("Missing required database environment variables")
	}
	// First, check if the database exists and create it if it doesn't
	serverDSN := "postgres://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/postgres?sslmode=verify-full"
	db, err := sql.Open("postgres", serverDSN)
	if err != nil {
		log.Fatal("failed to connect to MySQL server: ", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatal("failed to create database: ", err)
	}

	// Now, connect to the newly created or existing database
	dsn := "postgresql://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=verify-full"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	//Migrate the schema
	if err := DB.AutoMigrate(
		&models.ProjectAdmin{},
		&models.ProjectWork{},
		&models.ProjectUploader{},
	); err != nil {
		log.Fatal(err)
	}
}
