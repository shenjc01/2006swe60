package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"net/http"
)

func generateClientKey(sessionID string) error {
	// Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	// Convert keys to PEM format
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	spkiBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: spkiBytes,
	})

	// Store keys in the database
	_, err = DB.Exec("INSERT INTO SessionKeys (sessionID, privateKey, publicKey) VALUES (?, ?, ?)",
		sessionID, privateKeyPEM, publicKeyPEM)
	return err
}

func ServeClientPublicKey(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("sessionID")
	if sessionID == "" {
		http.Error(w, "Missing sessionID", http.StatusBadRequest)
		return
	}

	var publicKey string
	err := DB.QueryRow("SELECT publicKey FROM SessionKeys WHERE sessionID = (?)", sessionID).Scan(&publicKey)
	if err != nil {
		err := generateClientKey(sessionID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = DB.QueryRow("SELECT publicKey FROM SessionKeys WHERE sessionID = (?)", sessionID).Scan(&publicKey)
	}

	w.Header().Set("Content-Type", "application/x-pem-file")
	w.Write([]byte(publicKey))
	fmt.Printf("%x", publicKey)
}

// Store the decrypted AES key in the database for the client
func storeAESKey(clientID string, aesKey []byte) error {
	// Store the AES key in hex format
	aesKeyHex := hex.EncodeToString(aesKey)

	_, err := DB.Exec("UPDATE SessionKeys SET aesKey = (?) WHERE sessionID = (?)", aesKeyHex, clientID)
	return err
}

// Decrypt AES key received from client and store it in the database
func DecryptClientAESKey(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("sessionID")
	if clientID == "" {
		http.Error(w, "Missing clientID", http.StatusBadRequest)
		return
	}

	// Retrieve the private key from the database
	var privateKeyPEM string
	err := DB.QueryRow("SELECT privateKey FROM SessionKeys WHERE sessionID = (?)", clientID).Scan(&privateKeyPEM)
	if err != nil {
		http.Error(w, "SessionID not found", http.StatusNotFound)
		return
	}

	// Decode PEM-encoded private key
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		http.Error(w, "Invalid private key format", http.StatusInternalServerError)
		return
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		http.Error(w, "Failed to parse private key", http.StatusInternalServerError)
		return
	}

	// Read the encrypted AES key from the request body

	encryptedAESKey, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Failed to read encrypted AES key: %v\n", err)
		http.Error(w, "Failed to read encrypted AES key", http.StatusBadRequest)
		return
	}
	// Convert to base64 string
	base64String := base64.StdEncoding.EncodeToString(encryptedAESKey)
	fmt.Printf("Content-Length: %d\n", r.ContentLength)
	fmt.Printf("Actual-Length: %d\n", len(encryptedAESKey))
	fmt.Printf("Request Headers: %+v\n", r.Header)
	fmt.Printf("Key: %s\n", base64String)

	// Decrypt the AES key using the client's private key
	hash := sha256.New()
	aesKey, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, encryptedAESKey, nil)
	if err != nil {
		fmt.Printf("Failed to decrypt AES key. Error: %v\nEncrypted Key: %x\n", err, encryptedAESKey)
		http.Error(w, "Failed to decrypt AES key", http.StatusInternalServerError)
		return
	}

	// Store the AES key in the database
	if err := storeAESKey(clientID, aesKey); err != nil {
		fmt.Printf("Failed to store AES key. Error: %v\n", err)
		http.Error(w, "Failed to store AES key", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "AES key stored successfully for client: %s", clientID)
}

// Retrieve the AES key for a client
func getAESKey(clientID string) ([]byte, error) {
	var aesKeyHex string
	err := DB.QueryRow("SELECT aeskey FROM SessionKeys WHERE sessionID = (?)", clientID).Scan(&aesKeyHex)
	if err != nil {

		return nil, err
	}

	// Decode the hex format AES key to bytes
	aesKey, err := hex.DecodeString(aesKeyHex)
	return aesKey, err
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
	// Respond with a success message
	w.Write([]byte("Data received successfully"))

	key, err := getAESKey(req.SessionID)
	if err != nil {
		fmt.Printf("Failed to retrieve AES key. Error: %v\n", err)
		return
	}
	// Step 1: Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("failed to create AES cipher: %v", err)
	}

	// Step 2: Create a GCM cipher mode (Galois/Counter Mode)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Printf("failed to create GCM mode: %v", err)
	}

	// Step 3: Decrypt the ciphertext using the GCM cipher and the IV
	// gcm.Open expects the IV, ciphertext, and additional data (nil if not used)
	plaintext, err := gcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		fmt.Printf("failed to decrypt: %v", err)
	}

	// Step 2: Convert the byte array back to the original string
	originalString := string(plaintext)

	fmt.Printf("Original string: %s\n", originalString)
}
