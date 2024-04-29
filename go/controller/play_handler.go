package controller

import (
    "encoding/json"
    "fmt"
    "net/http"
    "double_up/utils"
)

type PlayStartRequest struct {
    PlayerID string `json:"playerId"`
    BetSize  int    `json:"betSize"`
}

type PlayResponse struct {
    CardDrawn int `json:"cardDrawn"`
}

func PlayStartHandler(w http.ResponseWriter, r *http.Request) {
    // Check if the request method is POST
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Decode the JSON request body into PlayStartRequest object
    var req PlayStartRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Print playerId and betSize to the console
    fmt.Println("Player ID:", req.PlayerID)
    fmt.Println("Bet Size:", req.BetSize)

    // Dummy logic: Generate a random number as the card drawn
    cardDrawn := 42 // Replace with your actual logic

    utils.Query()

    // Create a PlayResponse object
    response := PlayResponse{
        CardDrawn: cardDrawn,
    }

    // Set Content-Type header to application/json
    w.Header().Set("Content-Type", "application/json")

    // Encode the response object to JSON and send it in the response body
    json.NewEncoder(w).Encode(response)
}
