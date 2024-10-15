package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Location struct to represent data from the SQLite table
type Location struct {
	Name         string  `json:"Name"`
	OpeningHours string  `json:"Opening Hours"`
	Address      string  `json:"Address"`
	Longitude    float64 `json:"Longitude"`
	Latitude     float64 `json:"Latitude"`
}

type Comment struct {
	Longitude float64 `json:"Longitude"`
	Latitude  float64 `json:"Latitude"`
	Username  string  `json:"Username"`
	Comment   string  `json:"Comment"`
	Date      string  `json:"Date"`
}

// GetLocation handles GET requests to /api/location
func GetLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Example query to get all locations
	// Extract the category parameter from the query
	category := r.URL.Query().Get("category")

	// Check if the category is provided
	if category == "" {
		http.Error(w, "Category parameter is required", http.StatusBadRequest)
		return
	}

	// Use a parameterized query to safely insert the category
	rows, err := DB.Query(
		`SELECT Locations.*
		FROM Locations
		JOIN RecycleCategory 
		ON Locations.Latitude = RecycleCategory.Latitude 
		AND Locations.Longitude = RecycleCategory.Longitude
		WHERE RecycleCategory.RecycleItemCategory = ?`, category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to hold the results
	var locations []Location
	for rows.Next() {
		var loc Location
		if err := rows.Scan(&loc.Name, &loc.OpeningHours, &loc.Address, &loc.Latitude, &loc.Longitude); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		locations = append(locations, loc)
	}

	// Send the results as JSON
	json.NewEncoder(w).Encode(locations)
}

func GetLocationComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Example query to get all locations
	// Extract the category parameter from the query
	lat := r.URL.Query().Get("lat")
	long := r.URL.Query().Get("long")

	// Check if the category is provided
	if lat == "" || long == "" {
		http.Error(w, "Coordinate parameters required", http.StatusBadRequest)
		return
	}

	// Use a parameterized query to safely insert the category
	rows, err := DB.Query(
		`SELECT Comments.*
		FROM Comments
		WHERE Comments.Latitude = ? AND Comments.Longitude =?`, lat, long)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to hold the results
	var comments []Comment
	for rows.Next() {
		var loc Comment
		if err := rows.Scan(&loc.Latitude, &loc.Longitude, &loc.Username, &loc.Comment, &loc.Date); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		comments = append(comments, loc)
	}

	// Send the results as JSON
	json.NewEncoder(w).Encode(comments)
}

func AttemptLogin(w http.ResponseWriter, r *http.Request) {
	type EncryptedDataRequest struct {
		SessionID  string `json:"sessionid"`
		Username   string `json:"username"`
		Ciphertext string `json:"ciphertext"`
		IV         string `json:"iv"`
	}
	var req EncryptedDataRequest

	// Parse the JSON body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
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

	// Step 2: Convert the byte array back to the original string
	password := plaintext
	username := req.Username
	fmt.Printf("Original string: %x\n", password)

	var hashedPassword string
	// SQL query to retrieve the hashed password for the given username
	query := `SELECT hashedPassword FROM Users WHERE username = ?`

	// Execute the query and scan the result into the hashedPassword variable
	err = DB.QueryRow(query, username).Scan(&hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("username %s not found", username) // Username doesn't exist
			return
		}
		fmt.Println(err) // Some other error occurred
		return
	}
	fmt.Printf("Hashed Password Retrieved: %s\n", hashedPassword)

	var salt string
	query = `SELECT salt FROM Users WHERE username = ?`
	err = DB.QueryRow(query, username).Scan(&salt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("username %s not found", username) // Username doesn't exist
			return
		}
		fmt.Println(err) // Some other error occurred
		return
	}
	fmt.Printf("Hashed Password Retrieved: %s\n", hashedPassword)
	hashedPasswordBytes, err := hex.DecodeString(hashedPassword)
	if err != nil {
		fmt.Printf("failed to decode hex string of password: %v", err)
		return
	}
	saltBytes, err := hex.DecodeString(salt)
	if err != nil {
		fmt.Printf("failed to decode hex string of salt: %v", err)
		return
	}
	err = bcrypt.CompareHashAndPassword(hashedPasswordBytes, append(saltBytes, password...))
	if err != nil {
		fmt.Printf("Password mismatch %v", err)
		http.Error(w, `Password mismatch`, http.StatusBadRequest)
		return
	}

	loginIDBytes := make([]byte, 16)
	_, err = rand.Read(loginIDBytes)
	if err != nil {
		fmt.Printf("failed to generate login ID: %v", err)
		return
	}
	loginID := hex.EncodeToString(loginIDBytes)

	_, err = DB.Exec(`
		INSERT INTO LoggedIn (Username, LoginID) 
		VALUES (?, ?)
		ON CONFLICT(username) DO UPDATE SET LoginID = ?`, username, loginID, loginID)
	if err != nil {
		fmt.Printf("Failed to insert session: %v", err)
	}

	cookie := &http.Cookie{
		Name:     "loginID",
		Value:    loginID,
		Path:     "/",
		HttpOnly: true,  // Protect from JavaScript access
		Secure:   false, // Only send over HTTPS
		MaxAge:   3600,  // Optional: set expiration in seconds (1 hour here)
	}
	http.SetCookie(w, cookie)
	fmt.Println("Cookie Set")

	// Respond with a success message
	w.Write([]byte("Data received successfully"))
}
