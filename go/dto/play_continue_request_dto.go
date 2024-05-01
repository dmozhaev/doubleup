package dto

import (
    "github.com/google/uuid"
    "double_up/enums"
)

type PlayContinueRequestDto struct {
    PlayerID uuid.UUID `json:"PlayerID"`
    Choice   enums.SmallLargeChoice `json:"Choice"`
}
