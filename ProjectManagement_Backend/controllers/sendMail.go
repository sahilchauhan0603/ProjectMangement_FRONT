package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sahilchauhan0603/backend/helper"
	"github.com/sahilchauhan0603/backend/models"
)

// SendEmail sends a mail to reset password
// @Summary sends a mail to reset password
// @Description sends a mail to reset password
// @Tags email
// @Accept json
// @Produce json
// @Param email body models.EmailRequest true "Email to send"
// @Success 201 {object} models.EmailRequest
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to send email"
// @Router /api/v1/sendEmail [post]
func SendEmail(w http.ResponseWriter, r *http.Request) {

	var emailReq models.EmailRequest
	err := json.NewDecoder(r.Body).Decode(&emailReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	email := emailReq.Email
	fmt.Printf("Received email: %s\n", email)

	err = helper.SendEmail(email)
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email received"))
}
