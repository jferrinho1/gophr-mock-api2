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

	  // workspace endpoints
  mux.Handle(
    "/api/workspace-basic-info",
    &h.DummyHandler{
      FilePath: "./schemas/workspace-basic-info.json",
      Status:   200,
    },
  ).Methods("GET")
  
  mux.Handle(
    "/api/registration-info",
    &h.DummyHandler{
      FilePath: "./schemas/registration-info.json",
      Status:   200,
    },
  ).Methods("GET")
  
  mux.Handle(
    "/api/workspace-address",
    &h.DummyHandler{
      FilePath: "./schemas/workspace-address.json",
      Status:   200,
    },
  ).Methods("GET")
  
  mux.Handle(
    "/api/workspace-contacts",
    &h.DummyHandler{
      FilePath: "./schemas/workspace-contacts.json",
      Status:   200,
    },
  ).Methods("GET")
  
  mux.Handle(
    "/api/workspace-basic-info",
    &h.DummyHandler{
      Status:   200,
    },
  ).Methods("POST")
  
  mux.Handle(
    "/api/registration-info",
    &h.DummyHandler{
      Status:   200,
    },
  ).Methods("POST")
  
  mux.Handle(
    "/api/workspace-address",
    &h.DummyHandler{
      Status:   200,
    },
  ).Methods("POST")
  
  mux.Handle(
    "/api/workspace-contacts",
    &h.DummyHandler{
      Status:   200,
    },
  ).Methods("POST")
}
