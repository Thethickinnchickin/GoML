// middleware/rbacmiddleware.go

package middleware

import (
    "net/http"
   // "fmt"


    "github.com/Thethickinnchickin/GoML/helpers"
)




// RBACMiddleware enforces role-based access control
func RBACMiddleware(requiredRole string, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Your JWT validation and user authentication logic here
        // Check if the JWT token is valid
        // If valid, call next.ServeHTTP(w, r) to proceed
        // If invalid, return an error response

        if requiredRole == "admin" {
            if !helpers.IsAdminTokenValid(r) { // Replace IsTokenValid with your JWT validation logic
                errorMessage := "Unauthorized: Invalid JWT token"
                http.Error(w, errorMessage, http.StatusUnauthorized)
                return
            }
        } else if requiredRole == "user" {
            if !helpers.IsTokenValid(r) { // Replace IsTokenValid with your JWT validation logic
                errorMessage := "Unauthorized: Invalid JWT token"
                http.Error(w, errorMessage, http.StatusUnauthorized)
                return
            }
        }



        // Token is valid, proceed to the next handler
        next.ServeHTTP(w, r)
    })

}
