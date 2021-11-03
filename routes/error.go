package routes

import (
	"github.com/gorilla/mux"
	h "github.com/jferrinho1/gophr-mock-api2/handlers"
)

func RegisterErrorRoutes(mux *mux.Router) {
	mux.Handle(
		"/api/error/401",
		&h.DummyErrorHandler{Status: 401},
	)

	mux.Handle(
		"/api/error/404",
		&h.DummyErrorHandler{Status: 404},
	)

	mux.Handle(
		"/api/error/422",
		&h.DummyErrorHandler{Status: 422},
	)

	mux.Handle(
		"/api/error/500",
		&h.DummyErrorHandler{Status: 500},
	)
}
