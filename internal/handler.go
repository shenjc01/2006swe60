package internal

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// RegisterUser handles user registration requests
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

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
	_, err = DB.Exec("INSERT INTO Users (username, hashedPassword, salt) VALUES (?, ?, ?)", username, hex.EncodeToString(hashedPassword), hex.EncodeToString(salt))
	if err != nil {
		log.Printf("SQL Insert Error: %v", err)
		http.Error(w, "Username already taken or SQL error: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("User registered successfully!"))
}
