package dao

import (
    "fmt"
    "database/sql"
    "double_up/model"
)

func CreateGame(db *sql.DB, game *model.Game) error {
    query := "INSERT INTO game (id, player_id, created_at, bet_size, player_choice, card_drawn, potential_profit, game_result) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
    _, err := db.Exec(query, game.ID, game.PlayerID, game.CreatedAt, game.BetSize, game.PlayerChoice, game.CardDrawn, game.PotentialProfit, game.GameResult)
    if err != nil {
        return fmt.Errorf("Game cannot be created, id: %s. Error: %s", game.ID, err)
    }
    return nil
}
