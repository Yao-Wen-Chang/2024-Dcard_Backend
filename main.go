package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Ad represents the structure of an advertisement
type Ad struct {
	Title      string     `json:"title"`
	StartAt    time.Time  `json:"startAt"`
	EndAt      time.Time  `json:"endAt"`
	Conditions Conditions `json:"conditions"`
}

// Conditions represents the conditions for displaying the ad
type Conditions struct {
	AgeStart  int      `json:"ageStart,omitempty"`
	AgeEnd    int      `json:"ageEnd,omitempty"`
	Gender    string   `json:"gender,omitempty"`
	Country   []string `json:"country,omitempty"`
	Platform  []string `json:"platform,omitempty"`
}

var ads []Ad

func main() {
	r := mux.NewRouter()
	// Admin API
	r.HandleFunc("/api/v1/ad", createAd).Methods("POST")

	// Public API
	r.HandleFunc("/api/v1/ad", listAds).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)


}

func createAd(w http.ResponseWriter, r *http.Request) {
	var newAd Ad
	err := json.NewDecoder(r.Body).Decode(&newAd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ads = append(ads, newAd)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Ad created successfully"))
}

func listAds(w http.ResponseWriter, r *http.Request) {
	var age, gender, country, platform string
	var offset, limit int

	query := r.URL.Query()
	if val := query.Get("age"); val != "" {
		age = val
	}
	if val := query.Get("gender"); val != "" {
		gender = val
	}
	if val := query.Get("country"); val != "" {
		country = val
	}
	if val := query.Get("platform"); val != "" {
		platform = val
	}
	if val := query.Get("offset"); val != "" {
		// Parse offset as an integer
		// Handle error if not a valid integer
	}
	if val := query.Get("limit"); val != "" {
		// Parse limit as an integer
		// Handle error if not a valid integer
	}

	var matchingAds []Ad
	now := time.Now()

	for _, ad := range ads {
		// Check if ad is active (within startAt and endAt)
		if now.After(ad.StartAt) && now.Before(ad.EndAt) {
			// Check conditions
			if checkConditions(ad.Conditions, age, gender, country, platform) {
				matchingAds = append(matchingAds, ad)
			}
		}
	}

	// Pagination logic here using offset and limit

	response := map[string][]Ad{"items": matchingAds}
	json.NewEncoder(w).Encode(response)
}

func checkConditions(conditions Conditions, age, gender, country, platform string) bool {
	if conditions.AgeStart > 0 && conditions.AgeEnd > 0 {
		// Check age condition
		// If age is not within the range, return false
	}

	if conditions.Gender != "" {
		// Check gender condition
		// If gender does not match, return false
	}

	if len(conditions.Country) > 0 {
		// Check country condition
		// If country is not in the list, return false
	}

	if len(conditions.Platform) > 0 {
		// Check platform condition
		// If platform is not in the list, return false
	}

	return true
}