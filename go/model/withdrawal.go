package model

import (
    "time"
    "github.com/google/uuid"
)

type Withdrawal struct {
    ID        uuid.UUID
    PlayerID  uuid.UUID
    CreatedAt time.Time
    Amount    int64
}

func NewWithdrawal(playerID uuid.UUID, amount int64) *Withdrawal {
    return &Withdrawal{
        ID:        uuid.New(),       // Assuming you're using the github.com/google/uuid package for UUID generation
        PlayerID:  playerID,
        CreatedAt: time.Now().UTC(), // Using UTC time zone
        Amount:    amount,
    }
}
