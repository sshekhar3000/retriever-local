package main

import "encoding/json"

type Event struct {
    ID             string            `json:"id"`
    Title          string          `json:"title"`
    Category       string          `json:"category"`
    AlternateTitles json.RawMessage `json:"alternate_titles"`
    Labels         json.RawMessage `json:"labels"`
    Country        string          `json:"country"`
    Entities       json.RawMessage `json:"entities"`
    StartLocal     string          `json:"start_local"`
    EndLocal       string          `json:"end_local"`
    Timezone       string          `json:"timezone"`
    Location       json.RawMessage `json:"location"`
    Geo            json.RawMessage `json:"geo"`
    ArtistName     string          `json:"artist_name"`
}

type ArtistRequest struct {
    Artist string `json:"artist"`
}