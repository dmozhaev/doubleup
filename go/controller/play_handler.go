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
