package middleware

import (
    "net/http"
    "encoding/json"
    "golang.org/x/crypto/bcrypt"

	"github.com/Thethickinnchickin/GoML/repositories"
	"github.com/Thethickinnchickin/GoML/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    // Parse JSON data from the request body
    var credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    err := json.NewDecoder(r.Body).Decode(&credentials)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }

	hashedPasswordStr := string(hashedPassword)

    // Store the hashed password in your database or wherever you need
    // hashedPassword is a byte slice that you can save in your database
	newUser := models.User{
		Username: credentials.Username,
		PasswordHash:  hashedPasswordStr,
	}

	repositories.CreateUser(newUser)


    // Respond with a success message or do whatever you need to do next
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Password hashed and stored successfully"))
}
