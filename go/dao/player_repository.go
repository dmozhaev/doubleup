package dao

import (
    "fmt"
    "database/sql"
    "github.com/google/uuid"
    "double_up/model"
)

type PlayerRepository interface {
    FindById(id uuid.UUID) (*model.Player, error)
    //Save(player *Player) error
}

type PostgreSQLPlayerRepository struct {
    DB *sql.DB
}

func FindById(db *sql.DB, id uuid.UUID) (*model.Player, error) {
    query := "SELECT id, name, money_in_play, account_balance FROM player WHERE id = $1"
    row := db.QueryRow(query, id)

    var player model.Player
    err := row.Scan(&player.ID, &player.Name, &player.MoneyInPlay, &player.AccountBalance)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("player not found")
        }
        return nil, err
    }

    return &player, nil
}

//func (r *PostgreSQLPlayerRepository) Save(player *Player) error {
    // Implement logic to save the player to the database
//}
