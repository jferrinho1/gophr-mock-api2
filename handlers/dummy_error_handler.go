package handlers

import (
	"net/http"
)

type DummyErrorHandler struct {
	Status int
}

func (h *DummyErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(h.Status)
}
