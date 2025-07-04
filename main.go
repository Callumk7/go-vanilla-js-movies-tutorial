package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"callumkloos.dev/reelingit/data"
	"callumkloos.dev/reelingit/handlers"
	"callumkloos.dev/reelingit/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	// ENV VARIABLES
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file was available")
	}

	dbConn := os.Getenv("DATABASE_URL")
	if dbConn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatalf("Failed to connect to the DB: %v", err)
	}
	defer db.Close()

	// Init data repository
	movieRepo, err := data.NewMovieRepository(db, logInstance)

	if err != nil {
		log.Fatal("Failed to init movierepo")
	}

	movieHandler := handlers.MovieHandler{}

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

	// Handler for static files
	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = "localhost:8080"
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("Server failed", err)
	}
}

// add a comment down here
