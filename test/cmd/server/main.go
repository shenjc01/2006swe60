package main

import (
	"log"
	"net/http"
	"test/internal"
)

func main() {
	internal.InitDB()
	defer internal.DB.Close() // Close the database connection when the server stops

	// Serve HTML file (main page)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, "./web/MapPage.html")
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, "./web/Login.html")
	})
	http.HandleFunc("/final", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, "./web/FinalPage.html")
	})

	// Set up API endpoint for data
	http.HandleFunc("/api/location", internal.GetLocation) // GET requests for location
	http.HandleFunc("/api/locationcomment/", internal.GetLocationComment)
	http.HandleFunc("/getkey", internal.ServeClientPublicKey)
	http.HandleFunc("/sendkey", internal.DecryptClientAESKey)

	// Start the server
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
} // Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
