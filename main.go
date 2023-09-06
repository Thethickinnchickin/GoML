package main

import (
    "net/http"
    "fmt"

    "github.com/gorilla/mux"
    "github.com/rs/cors"

    "github.com/Thethickinnchickin/GoML/middleware"
    "github.com/Thethickinnchickin/GoML/handler"
    "github.com/Thethickinnchickin/GoML/db"
)

func main() {
    db.InitDB()

    r := mux.NewRouter()

    // Use authentication middleware for protected routes
    r.Handle("/profile", middleware.AuthenticationMiddleware(http.HandlerFunc(handler.ProtectedHandler)))
    //r.HandleFunc("/profile", handler.ProtectedHandler)

    // Use RBAC middleware for routes with role-based access control
    r.Handle("/admin", middleware.RBACMiddleware("admin", http.HandlerFunc(handler.AdminHandler)))

    // Use Hash To store registers password
    r.HandleFunc("/register", handler.RegisterHandler).Methods("POST")

    //Function to get login token
    r.HandleFunc("/login", handler.LoginHandler).Methods("POST")

    //Function to get login token
    r.HandleFunc("/admin/login", handler.AdminLoginHandler).Methods("POST")

    // Define your API routes and handlers here
    http.HandleFunc("/", handler.HomeHandler)

    // Create a new CORS handler with allowed origins
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"*"},
    })

    // Wrap your router with the CORS middleware
    handlerWithCORS := c.Handler(r)

    // Print a message to the console indicating that the server is up and running
    fmt.Println("Server is up and running on :8080")

    // Start your server with the CORS-wrapped handler
    http.ListenAndServe(":8080", handlerWithCORS)
}
