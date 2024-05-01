package integration

import (
    "bytes"
    "database/sql"
    "net/http"
    "net/http/httptest"
    "double_up/controller"
)

func SendRequestPlayStart(db *sql.DB, method string, requestBody string) (*httptest.ResponseRecorder) {
    req := httptest.NewRequest(method, "/play/start", bytes.NewBufferString(requestBody))

    // Create a response recorder to record the response
    rr := httptest.NewRecorder()

    // Convert controller.PlayStartHandler to a http.HandlerFunc
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        controller.PlayStartHandler(db, w, r)
    })

    // Serve the request using the handler and record the response
    handler.ServeHTTP(rr, req)

    return rr
}

func SendRequestPlayContinue(db *sql.DB, method string, requestBody string) (*httptest.ResponseRecorder) {
    req := httptest.NewRequest(method, "/play/continue", bytes.NewBufferString(requestBody))

    // Create a response recorder to record the response
    rr := httptest.NewRecorder()

    // Convert controller.PlayContinueHandler to a http.HandlerFunc
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        controller.PlayContinueHandler(db, w, r)
    })

    // Serve the request using the handler and record the response
    handler.ServeHTTP(rr, req)

    return rr
}

func SendRequestWithdraw(db *sql.DB, method string, requestBody string) (*httptest.ResponseRecorder) {
    req := httptest.NewRequest(method, "/withdraw/withdrawmoney", bytes.NewBufferString(requestBody))

    // Create a response recorder to record the response
    rr := httptest.NewRecorder()

    // Convert controller.WithdrawHandler to a http.HandlerFunc
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        controller.WithdrawHandler(db, w, r)
    })

    // Serve the request using the handler and record the response
    handler.ServeHTTP(rr, req)

    return rr
}
