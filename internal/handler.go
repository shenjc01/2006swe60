package internal

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"encoding/json"
)

// RegisterUser handles user registration requests
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	
	// Parse the JSON body
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email string `json:"email"`
	}
	
	// Decode the JSON Body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "message": "Invalid request body",
        })
		return
	}
	
	username := req.Username
	password := req.Password
	email := req.Email

	salt := make([]byte, 4)
	_, err = rand.Read(salt)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword(append(salt, []byte(password)...), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Insert the new user into the database
	_, err = DB.Exec("INSERT INTO Users (username, hashedPassword, salt, email) VALUES (?, ?, ?, ?)", username, hex.EncodeToString(hashedPassword), hex.EncodeToString(salt), email)
	if err != nil {
		log.Printf("SQL Insert Error: %v", err)
		w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "message": "Username already taken or SQL error",
        })
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "success": true,
        "message": "User registered successfully!",
    })
}
