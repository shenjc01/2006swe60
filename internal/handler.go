package internal

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"encoding/json"
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
)

// RegisterUser handles user registration requests
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	
	// Parse the JSON body
	var req struct {
		SessionID  string `json:"sessionid"`
		Username   string `json:"username"`
		Ciphertext string `json:"ciphertext"`
		IV         string `json:"iv"`
		Email	   string `json:"email"`
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

	// Base64 decode the ciphertext and IV
	ciphertext, err := base64.StdEncoding.DecodeString(req.Ciphertext)
	if err != nil {
		http.Error(w, "Invalid base64 ciphertext", http.StatusBadRequest)
		return
	}

	iv, err := base64.StdEncoding.DecodeString(req.IV)
	if err != nil {
		http.Error(w, "Invalid base64 IV", http.StatusBadRequest)
		return
	}

	// Now you have the username, ciphertext (byte slice), and IV (byte slice)
	fmt.Printf("SessionID: %s\n", req.SessionID)
	fmt.Printf("Username: %s\n", req.Username)
	fmt.Printf("Ciphertext: %x\n", ciphertext) // Print ciphertext in hex
	fmt.Printf("IV: %x\n", iv)                 // Print IV in hex

	key, err := getAESKey(req.SessionID)
	if err != nil {
		fmt.Printf("Failed to retrieve AES key. Error: %v\n", err)
		return
	}
	// Step 1: Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("failed to create AES cipher: %v", err)
		return
	}

	// Step 2: Create a GCM cipher mode (Galois/Counter Mode)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Printf("failed to create GCM mode: %v", err)
		return
	}

	// Step 3: Decrypt the ciphertext using the GCM cipher and the IV
	// gcm.Open expects the IV, ciphertext, and additional data (nil if not used)
	plaintext, err := gcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		fmt.Printf("failed to decrypt: %v", err)
		return
	}
	
	username := req.Username
	password := string(plaintext)
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

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	type UserProfile struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	
	// Get the username from the session
	username := GetUser(w, r)
	if username == "" {
		http.Error(w, "Not signed in", http.StatusForbidden)
		return
	}
	
	// Retrieve user profile from the database
	var userProfile UserProfile
	err := DB.QueryRow(`SELECT username, email FROM Users WHERE username = ?`, username).Scan(&userProfile.Username, &userProfile.Email)
	if err != nil {
		http.Error(w, "Error retrieving user data", http.StatusInternalServerError)
		log.Printf("Error retrieving user data: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userProfile)
}