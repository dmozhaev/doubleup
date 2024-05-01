package integration

import (
    "bytes"
    "database/sql"
    "io/ioutil"
    "encoding/json"
    "github.com/stretchr/testify/assert"
    "github.com/google/uuid"
    "net/http"
    "net/http/httptest"
    "testing"
    "double_up/controller"
    "double_up/dao"
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

func CheckDbTableCounts(t *testing.T, db *sql.DB, accessCount int, auditCount int, gameCount int, withdrawalCount int) {
    var count int
    db.QueryRow("SELECT count(id) FROM access_log").Scan(&count)
    assert.Equal(t, accessCount, count)
    db.QueryRow("SELECT count(id) FROM audit_log").Scan(&count)
    assert.Equal(t, auditCount, count)
    db.QueryRow("SELECT count(id) FROM game").Scan(&count)
    assert.Equal(t, gameCount, count)
    db.QueryRow("SELECT count(id) FROM withdrawal").Scan(&count)
    assert.Equal(t, withdrawalCount, count)
}

func CheckDbPlayerTable(t *testing.T, db *sql.DB, moneyInPlay int, accountBalance int) {
    playerID, _ := uuid.Parse("01162f1f-0bd9-43fe-8032-fa9590ee0e7e")
    player, _ := dao.FindPlayerById(db, playerID)
    assert.Equal(t, moneyInPlay, int(player.MoneyInPlay))
    assert.Equal(t, accountBalance, int(player.AccountBalance))
}
