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

func WithdrawHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    // check if the request method is POST
    if r.Method != http.MethodPost {
        errorResponse := ErrorResponse{Error: "WithdrawHandler: Method not allowed"}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(errorResponse)
        return
    }

    // deserialize the JSON request body into WithdrawRequestDto object
    var requestDto dto.WithdrawRequestDto
    if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
        fmt.Println("WithdrawHandler: dto is of incorrect format")
        errorResponse := ErrorResponse{Error: err.Error()}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(errorResponse)
        return
    }

    // check rate limit
    service.CheckAccessAllowed(db, r.RemoteAddr, "/withdraw/withdrawmoney")

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
    dtoErr := validation.ValidateWithdrawRequest(requestDto, player)
	if dtoErr != nil {
		fmt.Println("ValidateWithdrawRequest errors: ", dtoErr)
		errorResponse := ErrorResponse{Error: dtoErr.Error()}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(errorResponse)
		return
	}

    // withdrawal logic
    responseDto, withdrawErr := service.Withdraw(db, player)
    if err != nil {
        fmt.Println("Error in WithdrawGame: ", withdrawErr)
		errorResponse := ErrorResponse{Error: withdrawErr.Error()}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(errorResponse)
        return
    }

    // api OK response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(responseDto)
}
