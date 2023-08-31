package main

import (
    "net/http"
    "fmt"

    "github.com/gorilla/mux"
    "github.com/Thethickinnchickin/GoML/middleware"
	"github.com/Thethickinnchickin/GoML/models"
)

var userStore = map[string]models.User{
    "admin": {
        Username: "admin",
        Password: "password", 
        Role: "admin"
    },
}

func main() {
    r := mux.NewRouter()

    // Use authentication middleware for protected routes
    r.Handle("/profile", middleware.AuthenticationMiddleware(http.HandlerFunc(middleware.ProtectedHandler)))
    
    // Use RBAC middleware for routes with role-based access control
    r.Handle("/admin", middleware.RBACMiddleware("admin", http.HandlerFunc(middleware.AdminHandler)))

	//Function to get login token
	r.HandleFunc("/login",middleware.LoginHandler).Methods("POST")


    // Define your API routes and handlers here

    http.Handle("/", r)

	// Print a message to the console indicating that the server is up and running
    fmt.Println("Server is up and running on :8080")

    // Start your server
    http.ListenAndServe(":8080", nil)
}