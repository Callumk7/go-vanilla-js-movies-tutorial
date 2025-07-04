package handlers

import (
	"encoding/json"
	"net/http"

	"callumkloos.dev/reelingit/data"
	"callumkloos.dev/reelingit/logger"
)

type MovieHandler struct {
	Storage data.MovieStorage
	Logger  *logger.Logger
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Storage.GetTopMovies()

	if err != nil {
		http.Error(w, "Internal Error", 500)
		h.Logger.Error("Could not get top movies", err)
	}

	h.writeJSONResponse(w, movies)
}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Storage.GetRandomMovies()

	if err != nil {
		http.Error(w, "Internal Error", 500)
		h.Logger.Error("Could not get top movies", err)
	}

	h.writeJSONResponse(w, movies)
}

func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.Logger.Error("Failed to encode JSON", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
