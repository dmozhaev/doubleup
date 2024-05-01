package integration_tests

import (
    "database/sql"
    "fmt"
    "strings"
    "testing"
    "github.com/stretchr/testify/assert"
    "double_up/dao"
    "double_up/integration"
    "double_up/utils"
)

func setupDbPlayStart(db *sql.DB) {
	queryFunc := func(db *sql.DB, tx *sql.Tx) error {
        query := `
            DELETE FROM game;
            DELETE FROM withdrawal;
            DELETE FROM audit_log;
            DELETE FROM access_log;
            UPDATE player SET money_in_play = 0, account_balance = 1000 where id = '01162f1f-0bd9-43fe-8032-fa9590ee0e7e';
        `
        _, err := db.Exec(query)
        if err != nil {
            return err
        }
        return nil
	}
    dao.RunInTransaction(db, queryFunc)
}

func TestPlayStartBasic(t *testing.T) {
    // db connect
    db, err := utils.Connect()
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()

    // setup db
    setupDbPlayStart(db)

    // GET not allowed
    rr := integration.SendRequestPlayStart(db, "GET", `{
        "PlayerID": "01162f1f-0bd9-43fe-8032-fa9590ee0e7e",
        "BetSize": 10,
        "Choice": "SMALL"
    }`)
    assert.Equal(t, 500, rr.Code)
    assert.Equal(t, `{"error":"PlayStartHandler: Method not allowed"}`, strings.TrimSpace(rr.Body.String()))

    // PUT not allowed
    rr = integration.SendRequestPlayStart(db, "PUT", `{
        "PlayerID": "01162f1f-0bd9-43fe-8032-fa9590ee0e7e",
        "BetSize": 10,
        "Choice": "SMALL"
    }`)
    assert.Equal(t, 500, rr.Code)
    assert.Equal(t, `{"error":"PlayStartHandler: Method not allowed"}`, strings.TrimSpace(rr.Body.String()))

    // invalid PlayerID
    rr = integration.SendRequestPlayStart(db, "POST", `{
        "PlayerID": "asdasdadad",
        "BetSize": 10,
        "Choice": "SMALL"
    }`)
    assert.Equal(t, 500, rr.Code)
    assert.Equal(t, `{"error":"PlayStartHandler: invalid UUID length: 10"}`, strings.TrimSpace(rr.Body.String()))

    // missing PlayerID
    rr = integration.SendRequestPlayStart(db, "POST", `{
        "BetSize": 10,
        "Choice": "SMALL"
    }`)
    assert.Equal(t, 500, rr.Code)
    assert.Equal(t, `{"error":"PlayStartHandler: Player not found, id: 00000000-0000-0000-0000-000000000000"}`, strings.TrimSpace(rr.Body.String()))

    // player does not exist in DB
    rr = integration.SendRequestPlayStart(db, "POST", `{
        "PlayerID": "9ff66fec-17c4-4594-aa03-d053fc036bad",
        "BetSize": 10,
        "Choice": "SMALL"
    }`)
    assert.Equal(t, 500, rr.Code)
    assert.Equal(t, `{"error":"PlayStartHandler: Player not found, id: 9ff66fec-17c4-4594-aa03-d053fc036bad"}`, strings.TrimSpace(rr.Body.String()))

    // invalid BetSize
    rr = integration.SendRequestPlayStart(db, "POST", `{
        "PlayerID": "01162f1f-0bd9-43fe-8032-fa9590ee0e7e",
        "BetSize": "some string",
        "Choice": "SMALL"
    }`)
    assert.Equal(t, 500, rr.Code)
    assert.Equal(t, `{"error":"PlayStartHandler: json: cannot unmarshal string into Go struct field PlayStartRequestDto.BetSize of type int64"}`, strings.TrimSpace(rr.Body.String()))

    // invalid Choice
    rr = integration.SendRequestPlayStart(db, "POST", `{
        "PlayerID": "01162f1f-0bd9-43fe-8032-fa9590ee0e7e",
        "BetSize": 10,
        "Choice": "IDONTKNOW"
    }`)
    assert.Equal(t, 500, rr.Code)
    assert.Equal(t, `{"error":"PlayStartHandler: PlayValidator: choice is invalid"}`, strings.TrimSpace(rr.Body.String()))
}
