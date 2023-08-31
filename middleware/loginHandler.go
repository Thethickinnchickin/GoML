package middleware

import (
    "net/http"
	"time"

    "github.com/dgrijalva/jwt-go"
	"github.com/Thethickinnchickin/GoML/models"
	"encoding/json"
)

var userStore = map[string]models.User{
    "admin": {
        Username: "admin",
        Password: "password", 
		Role: "admin",
    },
}




func LoginHandler(w http.ResponseWriter, r *http.Request) {
    // Parse JSON data from the request body
    var credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }




    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    // Check if the user exists and the password is correct
    user, ok := userStore[credentials.Username]
    if !ok || user.Password != credentials.Password {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Generate a JWT token
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = credentials.Username
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 7 days in hours
	claims["exp"] = expirationTime.Unix()
    // Add other claims as needed...

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
