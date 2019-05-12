package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func SecretCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func SecretGet(w http.ResponseWriter, r *http.Request) {
	secret := Secret{
		SecretText: "hello",
		SecretHash: "1231231",
		ViewTime:   4,
		ExpiredAt:  time.Now(),
	}

	json.NewEncoder(w).Encode(secret)
}
