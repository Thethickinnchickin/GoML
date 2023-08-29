
// main.go

package main

import (
    "net/http"

    "github.com/gorilla/mux"
    "github.com/Thethickinnchickin/GoML/middleware" // Import your middleware package
)


func main() {
    r := mux.NewRouter()

    // Use authentication middleware for protected routes
    r.Handle("/protected", middleware.AuthenticationMiddleware(http.HandlerFunc(protectedHandler)))
    
    // Use RBAC middleware for routes with role-based access control
    r.Handle("/admin", middleware.RBACMiddleware("admin", http.HandlerFunc(adminHandler)))

    // Define your API routes and handlers here

    http.Handle("/", r)

    // Start your server
    http.ListenAndServe(":8080", nil)
}