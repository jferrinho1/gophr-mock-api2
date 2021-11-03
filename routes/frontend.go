package routes

import (
	"github.com/gorilla/mux"
	h "github.com/jferrinho1/gophr-mock-api2/handlers"
)

func RegisterFrontendRoutes(mux *mux.Router) {
	mux.Handle(
		"/api/test",
		&h.DummyHandler{
			FilePath: "./schemas/test.json",
			Status:   201,
		},
	).Methods("GET")

	subRouter := mux.NewRoute().Subrouter()
	subRouter.Use(m.RequireFrontendVersionHeader)
	
	subRouter.Handle(
		"/api/test",
		&h.DummyHandler{
			FilePath: "./schemas/test.json",
			Status:   201,
		},
	).Methods("GET")
}
