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

// PlayStartHandler handles the HTTP request for starting a game.
//
// It expects a POST request with JSON data containing the following fields:
//   - playerId: The ID of the player starting the game in the UUID format.
//   - choice: The choice made by the player (small or large card).
//   - betSize: The size of the bet for the game (number).
//
// If successful, it responds with JSON data containing the following fields:
//   - cardDrawn: The randomly drawn card for the game [1...13].
//   - gameResult: The result of the game (W or L).
//   - moneyInPlay: The amount of money in play after the game (amount of the possible win).
//   - remainingBalance: The remaining balance of the player after the game.
//
// If there's an error decoding the request body or processing the game, it
// responds with an appropriate HTTP error status code and an error message.
func PlayStartHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    // check rate limit
    service.CheckAccessAllowed(db, r.RemoteAddr, "/play/start")

    // check if the request method is POST
    if r.Method != http.MethodPost {
        HandlerError(w, "PlayStartHandler: Method not allowed")
        return
    }

    // deserialize the JSON request body into PlayStartRequestDto object
    var requestDto dto.PlayStartRequestDto
    if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
        HandlerError(w, fmt.Sprintf("PlayStartHandler: %s", err.Error()))
        return
    }

    // player should exist in DB
    player, err := service.GetPlayer(db, requestDto.PlayerID)
    if err != nil {
        HandlerError(w, fmt.Sprintf("PlayStartHandler: %s", err.Error()))
        return
    }

    // validate request dto
    dtoErr := validation.ValidateStartRequest(requestDto, player)
	if dtoErr != nil {
        HandlerError(w, fmt.Sprintf("PlayStartHandler: %s", dtoErr.Error()))
		return
	}

    // start game logic
    responseDto, startGameErr := service.StartGame(db, player, requestDto.BetSize, requestDto.Choice)
    if startGameErr != nil {
        HandlerError(w, fmt.Sprintf("PlayStartHandler: %s", startGameErr.Error()))
        return
    }

    // api OK response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(responseDto)
}

// PlayContinueHandler handles the HTTP request for continuing a game.
//
// It expects a POST request with JSON data containing the following fields:
//   - playerId: The ID of the player starting the game in the UUID format.
//   - choice: The choice made by the player (small or large card).
//
// If successful, it responds with JSON data containing the following fields:
//   - cardDrawn: The randomly drawn card for the game [1...13].
//   - gameResult: The result of the game (W or L).
//   - moneyInPlay: The amount of money in play after the game (amount of the possible win).
//   - remainingBalance: The remaining balance of the player after the game.
//
// If there's an error decoding the request body or processing the game, it
// responds with an appropriate HTTP error status code and an error message.
func PlayContinueHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    // check rate limit
    service.CheckAccessAllowed(db, r.RemoteAddr, "/play/continue")

    // check if the request method is POST
    if r.Method != http.MethodPost {
        HandlerError(w, "PlayContinueHandler: Method not allowed")
        return
    }

    // deserialize the JSON request body into PlayContinueRequestDto object
    var requestDto dto.PlayContinueRequestDto
    if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
        HandlerError(w, fmt.Sprintf("PlayContinueHandler: %s", err.Error()))
        return
    }

    // player should exist in DB
    player, err := service.GetPlayer(db, requestDto.PlayerID)
    if err != nil {
        HandlerError(w, fmt.Sprintf("PlayContinueHandler: %s", err.Error()))
        return
    }

    // validate request dto
    dtoErr := validation.ValidateContinueRequest(requestDto, player)
	if dtoErr != nil {
        HandlerError(w, fmt.Sprintf("PlayContinueHandler: %s", dtoErr.Error()))
		return
	}

    // continue game logic
    responseDto, continueGameErr := service.ContinueGame(db, player, player.MoneyInPlay, requestDto.Choice)
    if continueGameErr != nil {
        HandlerError(w, fmt.Sprintf("PlayContinueHandler: %s", continueGameErr.Error()))
        return
    }

    // api OK response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(responseDto)
}
