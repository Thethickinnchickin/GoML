// middleware/authmiddleware.go

package middleware

import (
    "net/http"
    "encoding/json"

    "github.com/Thethickinnchickin/GoML/helpers"
    //"fmt"
    "github.com/dgrijalva/jwt-go"

)

// AuthenticationMiddleware handles user authentication
func AuthenticationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Your JWT validation and user authentication logic here
        // Check if the JWT token is valid
        // If valid, call next.ServeHTTP(w, r) to proceed
        // If invalid, return an error response
        tokenString := helpers.ExtractTokenFromRequest(r)
        if !helpers.IsTokenValid(r) { // Replace IsTokenValid with your JWT validation logic
            errorMessage := "Unauthorized: Invalid JWT token"
            http.Error(w, errorMessage, http.StatusUnauthorized)
            return
        }
        // Parse and validate the token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // You should replace "your-secret-key" with your actual secret key
            return []byte("secret1"), nil
        })

        if err != nil || !token.Valid {
            errorMessage := "Unauthorized: Invalid JWT token"
            http.Error(w, errorMessage, http.StatusUnauthorized)
            return
        }

        // Access token claims
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            errorMessage := "Unauthorized: Failed to parse claims"
            http.Error(w, errorMessage, http.StatusUnauthorized)
            return
        }
        
        // Check if "username" key exists in claims
        username, ok := claims["username"].(string)
        if !ok {
            errorMessage := "Unauthorized: Missing 'username' claim"
            http.Error(w, errorMessage, http.StatusUnauthorized)
            return
        }
        
        // Create a response data structure
        responseData := struct {
            Username string `json:"username"`
        }{
            Username: username,
        }

        // Encode the response data as JSON
        jsonData, err := json.Marshal(responseData)
        if err != nil {
            errorMessage := "Internal Server Error: Failed to encode response data"
            http.Error(w, errorMessage, http.StatusInternalServerError)
            return
        }

        // Set the content type header and write the JSON data to the response
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(jsonData)

        // Token is valid, proceed to the next handler
        next.ServeHTTP(w, r)
        })
}



