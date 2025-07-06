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

	// Init account repository for users
	acccountRepo, err := data.NewAccountRepository(db, logInstance)
	if err != nil {
		log.Fatal("Failed to init the account repository")
	}

	movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)
	accountHandler := handlers.NewAccountHandler(acccountRepo, logInstance)

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)
	http.HandleFunc("/api/movies/search", movieHandler.SearchMovies)
	http.HandleFunc("/api/movies/", movieHandler.GetMovie)
	http.HandleFunc("/api/genres", movieHandler.GetGenres)

	http.HandleFunc("/api/account/register", accountHandler.Register)
	http.HandleFunc("/api/account/authenticate", accountHandler.Authenticate)

	catchAllClientRoutesHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	}

	http.HandleFunc("/movies", catchAllClientRoutesHandler)
	http.HandleFunc("/movies/", catchAllClientRoutesHandler)

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
