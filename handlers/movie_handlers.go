package handlers

import (
	"encoding/json"
	"net/http"

	"callumkloos.dev/reelingit/models"
)

type MovieHandler struct {
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     181,
			Title:       "The Hacker",
			ReleaseYear: 2022,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting: []models.Actor{{
				ID:        1,
				FirstName: "Max",
				LastName:  "Power",
			}},
		},
		{
			ID:          2,
			TMDB_ID:     181,
			Title:       "The Hacker 2",
			ReleaseYear: 2022,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting: []models.Actor{{
				ID:        1,
				FirstName: "Max",
				LastName:  "Power",
			}},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		// TODO: Log error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     181,
			Title:       "The Hacker",
			ReleaseYear: 2022,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting: []models.Actor{{
				ID:        1,
				FirstName: "Max",
				LastName:  "Power",
			}},
		},
		{
			ID:          2,
			TMDB_ID:     181,
			Title:       "The Hacker 2",
			ReleaseYear: 2022,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting: []models.Actor{{
				ID:        1,
				FirstName: "Max",
				LastName:  "Power",
			}},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		// TODO: Log error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
