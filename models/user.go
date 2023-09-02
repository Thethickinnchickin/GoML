package models


type User struct {
    Username string
    PasswordHash string
    Role     *string // Add a field to store the user's role
}

