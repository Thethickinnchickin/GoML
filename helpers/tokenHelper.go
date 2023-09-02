package helpers

import (
    "net/http"
	"strings"
    "time"

    "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("secret1")
var jwtAdminSecret = []byte("admin-secret")

// IsTokenValid validates a JWT token
func IsTokenValid(r *http.Request) bool {
    // Extract the JWT token from the "Authorization" header
    authorizationHeader := r.Header.Get("Authorization")
    if authorizationHeader == "" {
        return false
    }

	

    // The token should be in the format "Bearer <token>"
    parts := strings.Split(authorizationHeader, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
		
        return false
    }

	

    tokenString := parts[1]

    // Parse the JWT token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil {
        return false
    }

    // Check if the token is valid
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Check the "exp" (expiration) claim
        expClaim, ok := claims["exp"].(float64)
        if !ok {
            return false // "exp" claim is not a float64
        }

        expirationTime := time.Unix(int64(expClaim), 0)
        if time.Now().UTC().After(expirationTime) {
            return false // Token has expired
        }

        // Optionally, you can check other claims, such as issuer ("iss") and audience ("aud")

        // You can also access custom claims if needed
        role := claims["role"].(string)
        if role != "user" {
            return false
        }
        
        return true
    }

    return false
}

// IsTokenValid validates a JWT token
func IsAdminTokenValid(r *http.Request) bool {
    // Extract the JWT token from the "Authorization" header
    authorizationHeader := r.Header.Get("Authorization")
    if authorizationHeader == "" {
        return false
    }

	

    // The token should be in the format "Bearer <token>"
    parts := strings.Split(authorizationHeader, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
		
        return false
    }

	

    tokenString := parts[1]

    // Parse the JWT token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtAdminSecret, nil
    })

    if err != nil {
        return false
    }

    // Check if the token is valid
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Check the "exp" (expiration) claim
        expClaim, ok := claims["exp"].(float64)
        if !ok {
            return false // "exp" claim is not a float64
        }

        expirationTime := time.Unix(int64(expClaim), 0)
        if time.Now().UTC().After(expirationTime) {
            return false // Token has expired
        }

        // Optionally, you can check other claims, such as issuer ("iss") and audience ("aud")

        // You can also access custom claims if needed
        role := claims["role"].(string)
        if role != "admin" {
            return false
        }
        
        return true
    }

    return false
}
