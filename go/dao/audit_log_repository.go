package dao

import (
    "fmt"
    "database/sql"
    "double_up/model"
)

func CreateAuditLog(db *sql.DB, auditLog *model.AuditLog) error {
    query := "INSERT INTO audit_log (id, player_id, record_id, target_table, created_at, operation) VALUES ($1, $2, $3, $4, $5, $6)"
    _, err := db.Exec(query, auditLog.ID, auditLog.PlayerID, auditLog.RecordID, auditLog.TargetTable, auditLog.CreatedAt, auditLog.Operation)
    if err != nil {
        return fmt.Errorf("AuditLog cannot be created, id: %s. Error: %s", auditLog.ID, err)
    }
    return nil
}
