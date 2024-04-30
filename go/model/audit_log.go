package model

import (
	"time"
	"github.com/google/uuid"
	"double_up/enums"
)

type AuditLog struct {
	ID          uuid.UUID
	PlayerID    uuid.UUID
	RecordID    uuid.UUID
	TargetTable string
	CreatedAt   time.Time
	Operation   enums.AuditOperation
}

func NewAuditLog(playerID, recordID uuid.UUID, targetTable string, operation enums.AuditOperation) *AuditLog {
	return &AuditLog{
		ID:          uuid.New(),
		PlayerID:    playerID,
		RecordID:    recordID,
		TargetTable: targetTable,
		CreatedAt:   time.Now().UTC(),
		Operation:   operation,
	}
}
