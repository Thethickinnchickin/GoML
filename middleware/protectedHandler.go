// middleware/protectedHandler.go

package middleware

import (
    "net/http"
    "github.com/Thethickinnchickin/GoML/repositories"

    "fmt"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
    // Your code for the /protected route handler goes here
    // You can use w to write responses and r to read requests

    

    user, _ := repositories.GetUser("admin")
    fmt.Print(user);
}
