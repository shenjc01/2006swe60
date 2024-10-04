package internal

import (
	"encoding/json"
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
