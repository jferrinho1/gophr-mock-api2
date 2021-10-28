package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	// mux.Use(m.CheckAuthHeader)
	RegisterRoutes(mux)

	// Heroku gets port to bind on from $PORT env value
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	http.ListenAndServe(":"+port, mux)
}
