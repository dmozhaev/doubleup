package dao

import (
    "fmt"
    "database/sql"
    "github.com/google/uuid"
    "double_up/model"
)

func FindById(db *sql.DB, id uuid.UUID) (*model.Player, error) {
    query := "SELECT id, name, money_in_play, account_balance FROM player WHERE id = $1"
    row := db.QueryRow(query, id)

    var player model.Player
    err := row.Scan(&player.ID, &player.Name, &player.MoneyInPlay, &player.AccountBalance)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("Player not found, id: %s", id)
        }
        return nil, err
    }

    return &player, nil
}

func Save(db *sql.DB, player *model.Player) (*model.Player, error) {
    query := "UPDATE player SET money_in_play = $1, account_balance = $2 WHERE id = $3"
    _, err := db.Exec(query, player.MoneyInPlay, player.AccountBalance, player.ID)
    if err != nil {
        return nil, fmt.Errorf("Player cannot be saved, id: %s. Error: %s", player.ID, err)
    }
    return FindById(db, player.ID)
}
