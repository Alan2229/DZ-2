package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
)

type Input struct {
	InputString string `json:"inputString"`
}
type Output struct {
	OutputString string `json:"outputString"`
}

func Decode(w http.ResponseWriter, r *http.Request) {
	var input Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad request", http.StatusBadGateway)
		return
	}
	decode, err := base64.StdEncoding.DecodeString(input.InputString)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadGateway)
		return
	}
	output := Output{OutputString: string(decode)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&output)
}
