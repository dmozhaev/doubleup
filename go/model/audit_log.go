package model

import (
	"time"
	"github.com/google/uuid"
)

type AuditLog struct {
	ID          uuid.UUID
	PlayerID    uuid.UUID
	RecordID    uuid.UUID
	TargetTable string
	CreatedAt   time.Time
	Operation   AuditOperation
}

type AuditOperation string

const (
	Create AuditOperation = "CREATE"
	Update AuditOperation = "UPDATE"
	Delete AuditOperation = "DELETE"
)

func NewAuditLog(playerID, recordID uuid.UUID, targetTable string, operation AuditOperation) *AuditLog {
	return &AuditLog{
		ID:          uuid.New(),
		PlayerID:    playerID,
		RecordID:    recordID,
		TargetTable: targetTable,
		CreatedAt:   time.Now().UTC(),
		Operation:   operation,
	}
}
