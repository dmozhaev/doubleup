package dto

import (
    "github.com/google/uuid"
    "double_up/enums"
)

type PlayContinueRequestDto struct {
    PlayerID uuid.UUID
    Choice   enums.SmallLargeChoice
}
