package service

import (
    "fmt"
    "database/sql"
    "double_up/dao"
    "double_up/enums"
    "double_up/model"
)

func Withdraw(db *sql.DB, player *model.Player) (string, error) {
    withdrawal := model.NewWithdrawal(player.ID, player.MoneyInPlay)
    err := dao.CreateWithdrawal(db, withdrawal)
    if err != nil {
        fmt.Println(err)
        return "ERROR", err
    }
    WriteAuditLog(db, player, enums.Insert, withdrawal.ID, "withdrawal")

    player.AccountBalance += player.MoneyInPlay
    player.MoneyInPlay = 0
    _, err = dao.UpdatePlayer(db, player)
    if err != nil {
        fmt.Println(err)
        return "ERROR", err
    }
    WriteAuditLog(db, player, enums.Update, player.ID, "player")

    return "OK", nil
}
