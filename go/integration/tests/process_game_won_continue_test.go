package integration_tests

import (
    "database/sql"
    "fmt"
    "strings"
    "testing"
    "github.com/stretchr/testify/assert"
    "double_up/dao"
    "double_up/enums"
    "double_up/integration"
    "double_up/utils"
)

func setupDbProcessGameWonContinue(db *sql.DB) {
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

func TestProcessGameWonContinue(t *testing.T) {
    // db connect
    db, err := utils.Connect()
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()

    // setup db
    setupDbProcessGameWonContinue(db)

    // start play is not possible
    rr := integration.SendRequestPlayStart(db, "POST", `{
        "PlayerID": "01162f1f-0bd9-43fe-8032-fa9590ee0e7e",
        "BetSize": 10,
        "Choice": "SMALL"
    }`)
    assert.Equal(t, 500, rr.Code)
    assert.Equal(t, `{"error":"PlayStartHandler: PlayValidator: there should be no money in play in order to start!"}`, strings.TrimSpace(rr.Body.String()))
    integration.CheckDbTableCounts(t, db, 1, 1, 0, 0)
    integration.CheckDbPlayerTable(t, db, 20, 990)

    // continue play is possible
    rr = integration.SendRequestPlayContinue(db, "POST", `{
        "PlayerID": "01162f1f-0bd9-43fe-8032-fa9590ee0e7e",
        "Choice": "LARGE"
    }`)
    assert.Equal(t, 200, rr.Code)
    integration.CheckDbTableCounts(t, db, 2, 4, 1, 0)

    resp := integration.DeserializePlayResponse(rr)
    if resp.GameResult == enums.W {
        assert.Equal(t, enums.W, resp.GameResult)
        assert.Equal(t, int(40), int(resp.MoneyInPlay))
        integration.CheckDbPlayerTable(t, db, 40, 990)
    } else {
        assert.Equal(t, enums.L, resp.GameResult)
        assert.Equal(t, int(0), int(resp.MoneyInPlay))
        integration.CheckDbPlayerTable(t, db, 0, 990)
    }
    assert.Equal(t, int(990), int(resp.RemainingBalance))
}
