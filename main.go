package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/lib/pq"
	"github.com/gorilla/mux"
)


type Gender string
const (
	Male   Gender = "M"
	Female Gender = "F"
)

type Country string
const (
	Taiwan  Country = "TW"
	Japan  Country = "JP"
	
)

type Platform string
const (
	Android Platform = "android"
	IOS  Platform = "ios"
	Web Platform = "web"
)
// Conditions represents the conditions for displaying the ad
type Conditions struct {
	AgeStart  int      `json:"ageStart,omitempty"`
	AgeEnd    int      `json:"ageEnd,omitempty"`
	Gender    Gender   `json:"gender,omitempty"`
	Country   []Country `json:"country,omitempty"`
	Platform  []Platform `json:"platform,omitempty"`
}
// Ad represents the structure of an advertisement
type Ad struct {
	Title      string     `json:"title"`
	StartAt    time.Time  `json:"startAt"`
	EndAt      time.Time  `json:"endAt"`
	Conditions Conditions `json:"conditions"`
}

var ads []Ad
var db *sql.DB

func main() {
	r := mux.NewRouter()
	connectDB()
	// Admin API
	r.HandleFunc("/api/v1/ad", createAd).Methods("POST")

	// Public API
	r.HandleFunc("/api/v1/ad", listAds).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)


}

func connectDB() {
	var err error
	connStr := "user=postgres password=root dbname=addatabase sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func createAd(w http.ResponseWriter, r *http.Request) {
	var newAd Ad
	err := json.NewDecoder(r.Body).Decode(&newAd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ads = append(ads, newAd)
// fmt.Println(ads)
	if newAd.Conditions.Gender != "" {
		genderArray := pq.Array(newAd.Conditions.Gender)
		_, err = db.Exec("INSERT INTO advertisement (title, start_at, end_at, age_start, age_end, gender, country, platform) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		newAd.Title, newAd.StartAt, newAd.EndAt, newAd.Conditions.AgeStart, newAd.Conditions.AgeEnd, genderArray, pq.Array(newAd.Conditions.Country), pq.Array(newAd.Conditions.Platform))
	} else {
		_, err = db.Exec("INSERT INTO advekrtisement (title, start_at, end_at, age_start, age_end, country, platform) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		newAd.Title, newAd.StartAt, newAd.EndAt, newAd.Conditions.AgeStart, newAd.Conditions.AgeEnd, pq.Array(newAd.Conditions.Country), pq.Array(newAd.Conditions.Platform))
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusCreated)
	// w.Write([]byte("Ad created successfully"))
}

func listAds(w http.ResponseWriter, r *http.Request) {
	var age, gender, country, platform string
	// var offset, limit int

	query := r.URL.Query()
	if val := query.Get("age"); val != "" {
		age = val
	}
	if val := query.Get("now"); val != "" {
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
// Construct the SQL query based on the filters
	sqlQuery := "SELECT * FROM advertisement WHERE start_at <= NOW() AND end_at >= NOW()"
	if age != "" {
		sqlQuery += " AND age_start <= " + age + " AND age_end >= " + age
	}
	if gender != "" {
		sqlQuery += " AND gender = '" + gender + "'"
	} 

	if country != "" {
		sqlQuery += " AND '" + country + "' = ANY(country)"
	}
	if platform != "" {
		sqlQuery += " AND '" + platform + "' = ANY(country)"
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