package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type DummyHandler struct {
	FilePath string
	Status int
}

func (h *DummyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Open our jsonFile
	jsonFile, err := os.Open(h.FilePath)

		// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}

	json, err := ioutil.ReadAll(jsonFile)

	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if h.Status != 0 {
		w.WriteHeader(h.Status)
	}

	w.Write(json)
}
