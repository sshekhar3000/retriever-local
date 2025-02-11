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

func eventsByCountryDateHandler(w http.ResponseWriter, r *http.Request) {

    var req ArtistCountryDateRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    if condition := req.Artist == "" || req.Country == "" || req.Start == "" || req.End == ""; condition {
        http.Error(w, "Missing artist, country, start or end in request body", http.StatusBadRequest)
        return 
    }

    event, err := getRowsByCountryDate(req.Artist, req.Country, req.Start, req.End)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(event)
}

func eventsByCountry(w http.ResponseWriter, r *http.Request) {
    var req ArtistCountryRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    if req.Artist == "" || req.Country == "" {
        http.Error(w, "Missing artist or country in request body", http.StatusBadRequest)
        return
    }

    event, err := getRowsByCountry(req.Artist, req.Country)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(event)
}