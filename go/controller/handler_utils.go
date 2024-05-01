package controller

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type ErrorResponse struct {
    Error string `json:"error"`
}

func HandlerError(w http.ResponseWriter, err string) {
    fmt.Println(err)
    errorResponse := ErrorResponse{Error: err}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(errorResponse)
}
