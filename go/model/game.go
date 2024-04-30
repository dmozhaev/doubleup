package model

import (
    "time"
    "github.com/google/uuid"
    "double_up/enums"
)

type Game struct {
    ID              uuid.UUID
    PlayerID        uuid.UUID
    CreatedAt       time.Time
    BetSize         int64
    PlayerChoice    enums.SmallLargeChoice
    CardDrawn       int16
    PotentialProfit int64
    GameResult      enums.GameResult
}

func NewGame(playerID uuid.UUID, betSize int64, playerChoice enums.SmallLargeChoice, cardDrawn int16, potentialProfit int64, gameResult enums.GameResult) *Game {
    return &Game{
        ID:              uuid.New(),       // Assuming you're using the github.com/google/uuid package for UUID generation
        PlayerID:        playerID,
        CreatedAt:       time.Now().UTC(), // Using UTC time zone
        BetSize:         betSize,
        PlayerChoice:    playerChoice,
        CardDrawn:       cardDrawn,
        PotentialProfit: potentialProfit,
        GameResult:      gameResult,
    }
}
