package service

import (
    "database/sql"
    "github.com/google/uuid"
    "double_up/dao"
    "double_up/enums"
    "double_up/model"
)

func WriteAuditLog(db *sql.DB, player *model.Player, operation enums.AuditOperation, recordID uuid.UUID, targetTable string) error {
    auditLog := model.NewAuditLog(player.ID, recordID, targetTable, operation)
    return dao.CreateAuditLog(db, auditLog)
}
