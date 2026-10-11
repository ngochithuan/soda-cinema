package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"time"

	"github.com/joho/godotenv"
)

type Movie struct {
	ID               uint      `json:"id"`
	Title            string    `json:"title"`
	OriginalTitle    string    `json:"original_title"`
	Description      string    `json:"description"`
	Duration         int       `json:"duration"` // minute
	ReleaseDate      time.Time `json:"release_date"`
	EndDate          time.Time `json:"end_date"`
	AgeRating        string    `json:"age_rating"`
	Language         string    `json:"language"`
	SubtitleLanguage string    `json:"subtitle_language"`
	Country          string    `json:"country"`

	PosterURL  string `json:"poster_url"`
	BannerURL  string `json:"banner_url"`
	TrailerURL string `json:"trailer_url"`

	Status string `json:"status"` // COMING_SOON, NOW_SHOWING, AVALIBLE

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to soda cinema")
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)

	fmt.Printf("Server is running at %s\n", PORT)
	http.ListenAndServe(PORT, mux)

}
