// middleware/rbacmiddleware.go

package middleware

import (
    "net/http"
)

var userStore = map[string]models.User{
    "admin": {
        Username: "admin",
        Password: "password", 
        Role: "admin",
    },
}


// RBACMiddleware enforces role-based access control
func RBACMiddleware(requiredRole string, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Retrieve the user's role from the context or user model
        // Parse JSON data from the request body
        var credentials struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }

        userRole := getUserRoleFromContextOrModel(r) // Implement this function

        // Check if the user has the required role
        if userRole != requiredRole {
            http.Error(w, "Forbidden: Insufficient role", http.StatusForbidden)
            return
        }

        // If the user has the required role, proceed to the next handler
        next.ServeHTTP(w, r)
    })
}

func getUserRoleFromContextOrModel(r *http.Request) string {
    ctx := r.Context()
    userRole, ok := ctx.Value("userRole").(string)
    if ok {
        return userRole
    }

    return ""
}

