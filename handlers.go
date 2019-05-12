package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type secretPost struct {
	Secret           string `json:"secret"`
	ExpireAfterViews int    `json:"expireAfterViews"`
	ExpireAfter      int    `json:"expireAfter"`
}

func respondWithError(err error, w http.ResponseWriter, code int) {
	writeJSON(w, &JsonError{
		Message: err.Error(),
	}, code)
}

func SecretCreate(store *inMemoryStore) (http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decode := json.NewDecoder(r.Body)
		var data secretPost
		decode.Decode(data)

		secret, error := addSecret(data.Secret, data.ExpireAfterViews, data.ExpireAfter)

		if error != nil {
			respondWithError(error, w, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		error = json.NewEncoder(w).Encode(secret)

		if error != nil {
			respondWithError(error, w, http.StatusInternalServerError)
			return
		}
	}
}

func SecretGet(s *inMemoryStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		secret, error := s.readSecret(mux.Vars(r)["hash"])
		if error != nil {
			respondWithError(err, w, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		error = json.NewEncoder(w).Encode(secret)
		if error != nil {
			respondWithError(error, w, http.StatusInternalServerError)
			return
		}
	}
}
