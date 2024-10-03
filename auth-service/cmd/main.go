package main

import (
	"encoding/json"
	"net/http"
)

type AuthResponse struct {
	Message string `json:"message"`
}

const validToken = "valid_token_123"

func authenticate(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	if token == validToken {
		json.NewEncoder(w).Encode(AuthResponse{Message: "Authenticated"})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(AuthResponse{Message: "Unauthorized"})
	}
}

func main() {
	http.HandleFunc("/auth", authenticate)
	http.ListenAndServe(":8084", nil)
}
