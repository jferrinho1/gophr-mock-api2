package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type DummyHandler struct {
	FilePath string
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
	w.Write(json)
}
