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

type ErrorResponse struct {
    Error string `json:"error"`
}

// PlayStartHandler handles the HTTP request for starting a game.
//
// It expects a POST request with JSON data containing the following fields:
//   - playerId: The ID of the player starting the game in the UUID format.
//   - choice: The choice made by the player (small or large card).
//   - betSize: The size of the bet for the game.
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
    // check if the request method is POST
    if r.Method != http.MethodPost {
        errorResponse := ErrorResponse{Error: "PlayStartHandler: Method not allowed"}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(errorResponse)
        return
    }

    // deserialize the JSON request body into PlayStartRequestDto object
    var requestDto dto.PlayStartRequestDto
    if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
        fmt.Println("PlayStartHandler: dto is of incorrect format")
        errorResponse := ErrorResponse{Error: err.Error()}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(errorResponse)
        return
    }

    // check rate limit
    service.CheckAccessAllowed(db, r.RemoteAddr, "/play/start")

    // player should exist in DB
    player, err := service.GetPlayer(db, requestDto.PlayerID)
    if err != nil {
        fmt.Println("PlayerRepository.FindPlayerById: ", err)
        errorResponse := ErrorResponse{Error: err.Error()}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(errorResponse)
        return
    }

    // validate request dto
    dtoErr := validation.ValidateStartRequest(requestDto, player)
	if dtoErr != nil {
		fmt.Println("ValidateStartRequest errors: ", dtoErr)
		errorResponse := ErrorResponse{Error: dtoErr.Error()}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(errorResponse)
		return
	}

    // start game logic
    responseDto, startGameErr := service.StartGame(db, player, requestDto.BetSize, requestDto.Choice)
    if err != nil {
        fmt.Println("Error in StartGame: ", startGameErr)
		errorResponse := ErrorResponse{Error: startGameErr.Error()}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(errorResponse)
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
    // check if the request method is POST
    if r.Method != http.MethodPost {
        errorResponse := ErrorResponse{Error: "PlayContinueHandler: Method not allowed"}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(errorResponse)
        return
    }

    // deserialize the JSON request body into PlayContinueRequestDto object
    var requestDto dto.PlayContinueRequestDto
    if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
        fmt.Println("PlayContinueHandler: dto is of incorrect format")
        errorResponse := ErrorResponse{Error: err.Error()}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(errorResponse)
        return
    }

    // check rate limit
    service.CheckAccessAllowed(db, r.RemoteAddr, "/play/continue")

    // player should exist in DB
    player, err := service.GetPlayer(db, requestDto.PlayerID)
    if err != nil {
        fmt.Println("PlayerRepository.FindPlayerById: ", err)
        errorResponse := ErrorResponse{Error: err.Error()}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(errorResponse)
        return
    }

    // validate request dto
    dtoErr := validation.ValidateContinueRequest(requestDto, player)
	if dtoErr != nil {
		fmt.Println("ValidateContinueRequest errors: ", dtoErr)
		errorResponse := ErrorResponse{Error: dtoErr.Error()}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(errorResponse)
		return
	}

    // continue game logic
    responseDto, continueGameErr := service.ContinueGame(db, player, player.MoneyInPlay, requestDto.Choice)
    if err != nil {
        fmt.Println("Error in ContinueGame: ", continueGameErr)
		errorResponse := ErrorResponse{Error: continueGameErr.Error()}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(errorResponse)
        return
    }

    // api OK response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(responseDto)
}
