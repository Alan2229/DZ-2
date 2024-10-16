package handlers

import (
	"encoding/json"
	"net/http"
)

func Version(w http.ResponseWriter, r *http.Request) {
	r = r
	version := map[string]string{"version": "v:1.0.0"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(version)
}
