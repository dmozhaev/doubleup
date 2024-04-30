package service

import (
    "database/sql"
    "double_up/dao"
    "double_up/enums"
    "double_up/model"
)

func Withdraw(db *sql.DB, player *model.Player) (string, error) {
    withdrawal := model.NewWithdrawal(player.ID, player.MoneyInPlay)
    dao.CreateWithdrawal(db, withdrawal)
    WriteAuditLog(db, player, enums.Insert, withdrawal.ID, "withdrawal")

    player.AccountBalance += player.MoneyInPlay
    player.MoneyInPlay = 0
    dao.UpdatePlayer(db, player)
    WriteAuditLog(db, player, enums.Update, player.ID, "player")

    return "OK", nil
}
