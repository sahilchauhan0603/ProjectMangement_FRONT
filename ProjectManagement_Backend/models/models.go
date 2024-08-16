package models

import "time"

// Admin represents the Admin model
// @Description Admin model
type ProjectAdmin struct {

	// ID of the admin
	// required: true
	AdminID int64 `gorm:"primaryKey;autoIncrement"`

	// Name of the Admin
	// required: true
	AdminName string

	// UerName of the Admin
	// required: true
	Username string

	// Password of the Admin
	// required: true
	Password string

	ProjectIdsToVerify string

	// Uploader associated with the Admin
	Uploader []ProjectUploader `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Works associated with the Admin
	Work []ProjectWork `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Work represents the Work model
// @Description Work model
type ProjectWork struct {
	// WorkID is the primary key
	// required: true
	WorkID int64 `gorm:"primaryKey;autoIncrement"`

	// Title of the Work
	// required: true
	Title string

	// Type of the Work
	// required: true
	Type string

	// Approach used in the Work
	// required: true
	Approach string

	// Description of the Work
	// required: true
	Description string

	// Supervisor of the Work
	// required: true
	Supervisor string

	// AdminID is the foreign key to Admin
	// required: true
	AdminID int64 `gorm:"index"`

	// Proof related to the Work
	// required: true
	Proof string

	// Uploaders associated with the Work
	Uploaders []ProjectUploader `gorm:"foreignKey:WorkID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// CreatedAt timestamp
	StartDate time.Time

	EndDate time.Time

	Category string
}

// Uploader represents the uploader model
// @Description Uploader model
type ProjectUploader struct {
	// EnrollmentNo is the primary key
	// required: true
	EnrollmentNo int64 `gorm:"primaryKey;autoIncrement"`

	// FirstName of the uploader
	// required: true
	FirstName string

	// LastName of the uploader
	// required: true
	LastName string

	// InstitutionName of the uploader
	// required: true
	Department string

	// WorkID is the foreign key to Work
	// required: true
	WorkID int64 `gorm:"index"`

	// AdminID is the foreign key to Admin
	// required: true
	AdminID int64 `gorm:"index"`

	// Email of the uploader
	// required: true
	Email string `gorm:"unique"`

	PhoneNo string

	BatchStart time.Time

	BatchEnd time.Time

	GithubProfile string

	LinkedInProfile string

	FacultyName string

	Topic string

	ProjectUrl string

	Researchstart time.Time

	Researchend time.Time

	AboutResearch string

	AboutAppoarch string
}

// EmailRequest represents the EmailRequest model
// @Description EmailRequest model
type EmailRequest struct {

	// Email of the user
	// required: true
	Email string `json:"email"`
}
