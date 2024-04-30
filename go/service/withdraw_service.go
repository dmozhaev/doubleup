package service

import (
    "database/sql"
    "double_up/dao"
    "double_up/model"
)

func Withdraw(db *sql.DB, player *model.Player) (string, error) {
    player.AccountBalance += player.MoneyInPlay
    player.MoneyInPlay = 0
    dao.UpdatePlayer(db, player)

    return "OK", nil
}
