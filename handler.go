package main
import (
	"encoding/json"
	"net/http"
)

func eventsHandler(w http.ResponseWriter, r *http.Request) {
    var req ArtistRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    if req.Artist == "" {
        http.Error(w, "Missing artist in request body", http.StatusBadRequest)
        return
    }

    event, err := getRows(req.Artist)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(event)
}