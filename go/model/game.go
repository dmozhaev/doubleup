package model

import (
    "time"
    "github.com/google/uuid"
)

type Game struct {
    ID              uuid.UUID
    PlayerID        uuid.UUID
    CreatedAt       time.Time
    BetSize         int64
    PlayerChoice    SmallLargeChoice
    CardDrawn       int16
    PotentialProfit int64
    GameResult      GameResult
}

type SmallLargeChoice string

const (
    Small SmallLargeChoice = "SMALL"
    Large SmallLargeChoice = "LARGE"
)

type GameResult string

const (
    Win  GameResult = "WIN"
    Lose GameResult = "LOSE"
    Draw GameResult = "DRAW"
)

func NewGame(playerID uuid.UUID, betSize int64, playerChoice SmallLargeChoice, cardDrawn int16, potentialProfit int64, gameResult GameResult) *Game {
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
