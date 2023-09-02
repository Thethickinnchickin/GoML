// db/database.go
package db

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB // Change 'db' to 'DB' to make it accessible from other packages

func InitDB() {
    var err error
    DB, err = sql.Open("mysql", "mattreiley:password@tcp(localhost:1971)/goml") // Replace with your MySQL connection details
    if err != nil {
        log.Fatal(err)
    }

    // Test the connection
    if err = DB.Ping(); err != nil {
        log.Fatal(err)
    }
}