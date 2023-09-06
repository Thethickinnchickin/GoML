package handler

import (
    "net/http"
    "encoding/json"
    "golang.org/x/crypto/bcrypt"
	"time"
	"github.com/dgrijalva/jwt-go"


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

    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = newUser.Username
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 7 days in hours
	claims["exp"] = expirationTime.Unix()
    claims["role"] = "user"


    // Sign the token with a secret key
    tokenString, err := token.SignedString([]byte("secret1")) // Replace with your actual secret key
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }


    // Respond with the token
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"token": "` + tokenString + `"}`))
}
