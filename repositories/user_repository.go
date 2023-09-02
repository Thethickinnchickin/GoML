// repositories/user_repository.go
package repositories

import (
    "log"
	"database/sql"

    "github.com/Thethickinnchickin/GoML/models"
	"github.com/Thethickinnchickin/GoML/db"
)

func CreateUser(user models.User) error {
    _, err := db.DB.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", user.Username, user.PasswordHash)
    if err != nil {
        log.Println("Error creating user:", err)
        return err
    }
    return nil
}

func GetUser(username string) (*models.User, error) {
    // Execute a SELECT query to retrieve the user by username
    row := db.DB.QueryRow("SELECT username, password_hash FROM users WHERE username = ?", username)

    var user models.User
    err := row.Scan(&user.Username, &user.PasswordHash)
    if err != nil {
        if err == sql.ErrNoRows {
            // User not found, handle this case accordingly
            return nil, nil
        }
        log.Println("Error retrieving user:", err)
        return nil, err
    }

    return &user, nil
}


