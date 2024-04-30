package dao

import (
    "fmt"
    "database/sql"
    "double_up/model"
)

func CreateWithdrawal(db *sql.DB, withdrawal *model.Withdrawal) error {
    query := "INSERT INTO withdrawal (id, player_id, created_at, amount) VALUES ($1, $2, $3, $4)"
    _, err := db.Exec(query, withdrawal.ID, withdrawal.PlayerID, withdrawal.CreatedAt, withdrawal.Amount)
    if err != nil {
        return fmt.Errorf("Withdrawal cannot be created, id: %s. Error: %s", withdrawal.ID, err)
    }
    return nil
}
