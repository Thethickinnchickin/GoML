// middleware/authmiddleware.go

package middleware

import (
    "net/http"

    "github.com/Thethickinnchickin/GoML/helpers"
)

// AuthenticationMiddleware handles user authentication
func AuthenticationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Your JWT validation and user authentication logic here
        // Check if the JWT token is valid
        // If valid, call next.ServeHTTP(w, r) to proceed
        // If invalid, return an error response

        if !helpers.IsTokenValid(r) { // Replace IsTokenValid with your JWT validation logic
            errorMessage := "Unauthorized: Invalid JWT token"
            http.Error(w, errorMessage, http.StatusUnauthorized)
            return
        }

        // Token is valid, proceed to the next handler
        next.ServeHTTP(w, r)
    })
}



