package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JoStMc/kundokubungo/internal/models"
	"github.com/joho/godotenv"
)

var sentenceStore models.Sentence

func main() {
	godotenv.Load(".env")


	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM environment variable is not set")
	}

	filepathRoot := os.Getenv("FILEPATH_ROOT")
	if filepathRoot == "" {
		log.Fatal("FILEPATH_ROOT environment variable is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	mux := http.NewServeMux()
	appHandler := http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))

	mux.Handle("/app/", appHandler)
	mux.HandleFunc("POST /api/sentences", handlerCreate)
	mux.HandleFunc("PATCH /api/sentences/{id}", handlerUpdate)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on: http://localhost:%s/app/\n", port)
	log.Fatal(srv.ListenAndServe())
}
