package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"

	h "github.com/jferrinho1/gophr-mock-api2/handlers"
	m "github.com/jferrinho1/gophr-mock-api2/middleware"
	r "github.com/jferrinho1/gophr-mock-api2/routes"
)

func main() {
	mux := mux.NewRouter()

	r.RegisterErrorRoutes(mux)
	registerOptionsRoute(mux)
	registerFrontendRoutes(mux)
	registerMobileRoutes(mux)

	// Heroku gets port to bind on from $PORT env value
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	http.ListenAndServe(":"+port, mux)
}

func registerFrontendRoutes(mux *mux.Router) {
	subRouter := mux.PathPrefix("/frontend").Subrouter()

	r.RegisterFrontendRoutes(subRouter)
	m.RegisterFrontendMiddleware(subRouter)
}

func registerMobileRoutes(mux *mux.Router) {
	subRouter := mux.PathPrefix("/mobile").Subrouter()

	r.RegisterMobileRoutes(subRouter)
	m.RegisterMobileMiddleware(subRouter)
}

func registerOptionsRoute(mux *mux.Router) {
	mux.Use(m.AddCorsHeaders)
	mux.PathPrefix("/").Handler(&h.DummyHandler{Status: 200}).Methods(http.MethodOptions)
}
