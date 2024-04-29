package dto

type PlayContinueRequestDto struct {
    PlayerID uuid.UUID
    Choice   SmallLargeChoice
}
