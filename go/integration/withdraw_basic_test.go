package integration

import (
    "bytes"
    "database/sql"
    "net/http"
    "net/http/httptest"
    "fmt"
    "strings"
    "testing"
    "github.com/stretchr/testify/assert"
    "double_up/controller"
    "double_up/dao"
    "double_up/utils"
)

func setupDbWithdraw(db *sql.DB) {
	queryFunc := func(db *sql.DB, tx *sql.Tx) error {
        query := `
            DELETE FROM game;
            DELETE FROM withdrawal;
            DELETE FROM audit_log;
            DELETE FROM access_log;
            UPDATE player SET money_in_play = 20, account_balance = 990 where id = '01162f1f-0bd9-43fe-8032-fa9590ee0e7e';
        `
        _, err := db.Exec(query)
        if err != nil {
            return err
        }
        return nil
	}
    dao.RunInTransaction(db, queryFunc)
}

func sendRequestWithdraw(db *sql.DB, method string, url string, requestBody string) (*httptest.ResponseRecorder) {
    req := httptest.NewRequest(method, url, bytes.NewBufferString(requestBody))

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

func TestWithdrawBasic(t *testing.T) {
    // db connect
    db, err := utils.Connect()
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()

    // setup db
    setupDbWithdraw(db)

    // GET not allowed
    rr := sendRequestWithdraw(db, "GET", "/withdraw/withdrawmoney", `{
        "PlayerID": "01162f1f-0bd9-43fe-8032-fa9590ee0e7e"
    }`)
    assert.Equal(t, 500, rr.Code)
    assert.Equal(t, `{"error":"WithdrawHandler: Method not allowed"}`, strings.TrimSpace(rr.Body.String()))

    // PUT not allowed
    rr = sendRequestWithdraw(db, "PUT", "/withdraw/withdrawmoney", `{
        "PlayerID": "01162f1f-0bd9-43fe-8032-fa9590ee0e7e"
    }`)
    assert.Equal(t, 500, rr.Code)
    assert.Equal(t, `{"error":"WithdrawHandler: Method not allowed"}`, strings.TrimSpace(rr.Body.String()))

    // invalid PlayerID
    rr = sendRequestWithdraw(db, "POST", "/withdraw/withdrawmoney", `{
        "PlayerID": "asdasdadad"
    }`)
    assert.Equal(t, 500, rr.Code)
    assert.Equal(t, `{"error":"WithdrawHandler: invalid UUID length: 10"}`, strings.TrimSpace(rr.Body.String()))

    // missing PlayerID
    rr = sendRequestWithdraw(db, "POST", "/withdraw/withdrawmoney", `{}`)
    assert.Equal(t, 500, rr.Code)
    assert.Equal(t, `{"error":"WithdrawHandler: Player not found, id: 00000000-0000-0000-0000-000000000000"}`, strings.TrimSpace(rr.Body.String()))

    // player does not exist in DB
    rr = sendRequestWithdraw(db, "POST", "/withdraw/withdrawmoney", `{
        "PlayerID": "9ff66fec-17c4-4594-aa03-d053fc036bad"
    }`)
    assert.Equal(t, 500, rr.Code)
    assert.Equal(t, `{"error":"WithdrawHandler: Player not found, id: 9ff66fec-17c4-4594-aa03-d053fc036bad"}`, strings.TrimSpace(rr.Body.String()))
}
