package integration

import (
    "bytes"
    "database/sql"
    "io/ioutil"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "double_up/controller"
    "double_up/dto"
)

func sendRequest(db *sql.DB, method string, url string, requestBody string, handlerFunc func(dbParam *sql.DB, wParam http.ResponseWriter, rParam *http.Request)) (*httptest.ResponseRecorder) {
    req := httptest.NewRequest(method, url, bytes.NewBufferString(requestBody))

    // Create a response recorder to record the response
    rr := httptest.NewRecorder()

    // Convert controller.PlayStartHandler to a http.HandlerFunc
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        handlerFunc(db, w, r)
    })

    // Serve the request using the handler and record the response
    handler.ServeHTTP(rr, req)

    return rr
}

func SendRequestPlayStart(db *sql.DB, method string, requestBody string) (*httptest.ResponseRecorder) {
    return sendRequest(db, method, "/play/start", requestBody, controller.PlayStartHandler)
}

func SendRequestPlayContinue(db *sql.DB, method string, requestBody string) (*httptest.ResponseRecorder) {
    return sendRequest(db, method, "/play/continue", requestBody, controller.PlayContinueHandler)
}

func SendRequestWithdraw(db *sql.DB, method string, requestBody string) (*httptest.ResponseRecorder) {
    return sendRequest(db, method, "/withdraw/withdrawmoney", requestBody, controller.WithdrawHandler)
}

func DeserializePlayResponse(rr *httptest.ResponseRecorder) (dto.PlayResponseDto) {
    body, _ := ioutil.ReadAll(rr.Body)
    var data dto.PlayResponseDto
    json.Unmarshal(body, &data)
    return data
}
