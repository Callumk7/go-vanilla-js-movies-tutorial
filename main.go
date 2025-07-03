package main

import (
	"log"
	"net/http"

	"callumkloos.dev/reelingit/handlers"
	"callumkloos.dev/reelingit/logger"
)

func initLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie.log")

	if err != nil {
		log.Fatalf("failed to init logger %v", err)
	}

	defer logInstance.Close()
	return logInstance

}

func main() {

	logInstance := initLogger()

	movieHandler := handlers.MovieHandler{}

	// Handler for static files
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/movies/top", movieHandler.GetTopMovies)

	const addr = "localhost:8080"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("Server failed", err)
	}
}

// add a comment down here
