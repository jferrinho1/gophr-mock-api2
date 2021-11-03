package handlers

import (
	"io/ioutil"
	"net/http"
	"os"
)

type DummyHandler struct {
	FilePath string
	Status int
}

func (h *DummyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.Status == 0 {
		h.Status = http.StatusOK
	}

	w.WriteHeader(h.Status)
	
	if h.FilePath != "" {
		jsonFile, _ := os.Open(h.FilePath)
		json, _ := ioutil.ReadAll(jsonFile)
		w.Write(json)
	}
}
