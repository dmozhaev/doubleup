package model

import (
	"github.com/google/uuid"
)

type Player struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	Name          string
	MoneyInPlay   int64
	AccountBalance int64
}
