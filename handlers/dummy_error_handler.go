package handlers

import (
	"net/http"
)

type DummyErrorHandler struct {
	Status int
}

func (h *DummyErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(h.Status)
}
