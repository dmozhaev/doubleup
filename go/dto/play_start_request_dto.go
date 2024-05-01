package dto

import (
    "github.com/google/uuid"
    "double_up/enums"
)

type PlayStartRequestDto struct {
    PlayerID uuid.UUID `json:"PlayerID"`
    Choice   enums.SmallLargeChoice `json:"Choice"`
    BetSize int64 `json:"BetSize"`
}
