package routes

import (
	"github.com/gorilla/mux"
	h "github.com/jferrinho1/gophr-mock-api2/handlers"
)

func RegisterFrontendRoutes(mux *mux.Router) {
	mux.Handle(
		"/api/test",
		&h.DummyHandler{
			FilePath: "./schemas/frontend/test.json",
			Status:   201,
		},
	).Methods("GET")

	  // workspace get endpoints
  mux.Handle(
    "/api/workspace-basic-info",
    &h.DummyHandler{
      FilePath: "./schemas/frontend/workspace-basic-info.json",
      Status:   200,
    },
  ).Methods("GET")
  
  mux.Handle(
    "/api/registration-info",
    &h.DummyHandler{
      FilePath: "./schemas/frontend/registration-info.json",
      Status:   200,
    },
  ).Methods("GET")
  
  mux.Handle(
    "/api/workspace-address",
    &h.DummyHandler{
      FilePath: "./schemas/frontend/workspace-address.json",
      Status:   200,
    },
  ).Methods("GET")
  
  mux.Handle(
    "/api/workspace-contacts",
    &h.DummyHandler{
      FilePath: "./schemas/frontend/workspace-contacts.json",
      Status:   200,
    },
  ).Methods("GET")
  
	// workspace post endpoints
  mux.Handle(
    "/api/workspace-basic-info",
    &h.DummyHandler{
			FilePath: "./schemas/frontend/standard-schema.json",
      Status:   200,
    },
  ).Methods("POST")
  
  mux.Handle(
    "/api/registration-info",
    &h.DummyHandler{
			FilePath: "./schemas/frontend/standard-schema.json",
      Status:   200,
    },
  ).Methods("POST")
  
  mux.Handle(
    "/api/workspace-address",
    &h.DummyHandler{
			FilePath: "./schemas/frontend/standard-schema.json",
      Status:   200,
    },
  ).Methods("POST")
  
  mux.Handle(
    "/api/workspace-contacts",
    &h.DummyHandler{
			FilePath: "./schemas/frontend/standard-schema.json",
      Status:   200,
    },
  ).Methods("POST")

	// billing overview get endpoints
 mux.Handle(
    "/api/billing-statement",
    &h.DummyHandler{
      FilePath: "./schemas/frontend/billing-overview-statement.json",
      Status:   200,
    },
  ).Methods("GET")
	
	// billing overview get endpoints
 mux.Handle(
    "/api/teams",
    &h.DummyHandler{
      FilePath: "./schemas/frontend/teams.json",
      Status:   200,
    },
  ).Methods("GET")

}
