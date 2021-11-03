package main

import (
	"net/http"
	"os"

	gm "github.com/gorilla/mux"

	m "github.com/jferrinho1/gophr-mock-api2/middleware"
	r "github.com/jferrinho1/gophr-mock-api2/routes"
)

func main() {
	mux := gm.NewRouter()

	r.RegisterErrorRoutes(mux)
	registerFrontendRoutes(mux)
	registerMobileRoutes(mux)

	// Heroku gets port to bind on from $PORT env value
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	http.ListenAndServe(":"+port, mux)
}

func registerFrontendRoutes(mux *gm.Router) {
	subRouter := mux.PathPrefix("/frontend").Subrouter()

	r.RegisterFrontendRoutes(subRouter)
	m.RegisterFrontendMiddleware(subRouter)
}

func registerMobileRoutes(mux *gm.Router) {
	subRouter := mux.PathPrefix("/mobile").Subrouter()

	r.RegisterMobileRoutes(subRouter)
	m.RegisterMobileMiddleware(subRouter)
}
