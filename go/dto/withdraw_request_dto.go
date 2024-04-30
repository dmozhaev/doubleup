package dto

import (
    "github.com/google/uuid"
)

type WithdrawRequestDto struct {
    PlayerID uuid.UUID
}
