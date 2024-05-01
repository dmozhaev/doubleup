package controller

import (
    "encoding/json"
    "fmt"
    "net/http"
    "database/sql"
    "double_up/service"
    "double_up/dto"
    "double_up/validation"
)

// PlayContinueHandler handles the HTTP request for continuing a game.
//
// It expects a POST request with JSON data containing the following fields:
//   - playerId: The ID of the player starting the game in the UUID format.
//
// If successful, it responds with the "OK" string.
//
// If there's an error decoding the request body or processing the game, it
// responds with an appropriate HTTP error status code and an error message.
func WithdrawHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    // check rate limit
    service.CheckAccessAllowed(db, r.RemoteAddr, "/withdraw/withdrawmoney")

    // check if the request method is POST
    if r.Method != http.MethodPost {
        HandlerError(w, "WithdrawHandler: Method not allowed")
        return
    }

    // deserialize the JSON request body into WithdrawRequestDto object
    var requestDto dto.WithdrawRequestDto
    if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
        HandlerError(w, fmt.Sprintf("WithdrawHandler: %s", err.Error()))
        return
    }


    // player should exist in DB
    player, err := service.GetPlayer(db, requestDto.PlayerID)
    if err != nil {
        HandlerError(w, fmt.Sprintf("WithdrawHandler: %s", err.Error()))
        return
    }

    // validate request dto
    dtoErr := validation.ValidateWithdrawRequest(requestDto, player)
	if dtoErr != nil {
        HandlerError(w, fmt.Sprintf("WithdrawHandler: %s", dtoErr.Error()))
		return
	}

    // withdrawal logic
    msg, withdrawErr := service.Withdraw(db, player)
    if err != nil {
        HandlerError(w, fmt.Sprintf("WithdrawHandler: %s", withdrawErr.Error()))
        return
    }

    // api OK response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(msg)
}
