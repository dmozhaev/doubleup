package controller

import (
    "encoding/json"
    "fmt"
    "net/http"
    "database/sql"
    "double_up/service"
    "double_up/dto"
)

type PlayResponse struct {
    CardDrawn int `json:"cardDrawn"`
}

func PlayStartHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    // check if the request method is POST
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // decode the JSON request body into PlayStartRequestDto object
    var requestDto dto.PlayStartRequestDto
    if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // player should exist in DB
    player, err := service.GetPlayer(db, requestDto.PlayerID)
    if err != nil {
        fmt.Println("Player with ID: %s no found in DB!", err)
        return
    }
fmt.Println("AccountBalance:", player.AccountBalance)



    // Print playerId and betSize to the console
    fmt.Println("Player ID:", requestDto.PlayerID)
    fmt.Println("Bet Size:", requestDto.BetSize)

    // Dummy logic: Generate a random number as the card drawn
    cardDrawn := 42 // Replace with your actual logic





    // Create a PlayResponse object
    response := PlayResponse{
        CardDrawn: cardDrawn,
    }

    // Set Content-Type header to application/json
    w.Header().Set("Content-Type", "application/json")

    // Encode the response object to JSON and send it in the response body
    json.NewEncoder(w).Encode(response)
}
