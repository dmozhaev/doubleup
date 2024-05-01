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

func setupDbProcessGameWonWithdraw(db *sql.DB) {
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

func TestProcessGameWonWithdraw(t *testing.T) {
    // db connect
    db, err := utils.Connect()
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()

    // setup db
    setupDbProcessGameWonWithdraw(db)

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

    // withdrawal is possible
    rr = integration.SendRequestWithdraw(db, "POST", `{
        "PlayerID": "01162f1f-0bd9-43fe-8032-fa9590ee0e7e"
    }`)
    assert.Equal(t, 200, rr.Code)
    assert.Equal(t, `"OK"`, strings.TrimSpace(rr.Body.String()))
    integration.CheckDbTableCounts(t, db, 2, 4, 0, 1)
    integration.CheckDbPlayerTable(t, db, 0, 1010)
}
