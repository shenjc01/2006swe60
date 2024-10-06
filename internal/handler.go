package internal

import (
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
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Insert the new user into the database
	_, err = DB.Exec("INSERT INTO Users (username, hashedPassword, email) VALUES (?, ?, ?)", username, string(hashedPassword), email)
	if err != nil {
		log.Printf("SQL Insert Error: %v", err)
		http.Error(w, "Username already taken or SQL error: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("User registered successfully!"))
}