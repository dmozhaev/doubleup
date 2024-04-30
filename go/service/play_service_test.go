package service

import (
	"double_up/enums"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestProcessGame(t *testing.T) {
    // player choice: Small, winning numbers: [1...6]
    for i := 1; i <= 6; i++ {
        result := ProcessGame(i, 10, enums.Small, 99)

        assert.Equal(t, int16(i), result.CardDrawn)
        assert.Equal(t, enums.W, result.GameResult)
        assert.Equal(t, int64(20), result.MoneyInPlay)
        assert.Equal(t, int64(99), result.RemainingBalance)
    }

	// player choice: Small, losing numbers: [7...13]
    for i := 7; i <= 13; i++ {
        result := ProcessGame(i, 10, enums.Small, 99)

        assert.Equal(t, int16(i), result.CardDrawn)
        assert.Equal(t, enums.L, result.GameResult)
        assert.Equal(t, int64(0), result.MoneyInPlay)
        assert.Equal(t, int64(99), result.RemainingBalance)
    }

    // player choice: Large, losing numbers: [1...7]
    for i := 1; i <= 7; i++ {
        result := ProcessGame(i, 10, enums.Large, 99)

        assert.Equal(t, int16(i), result.CardDrawn)
        assert.Equal(t, enums.L, result.GameResult)
        assert.Equal(t, int64(0), result.MoneyInPlay)
        assert.Equal(t, int64(99), result.RemainingBalance)
    }

	// player choice: Large, winning numbers: [8...13]
    for i := 8; i <= 13; i++ {
        result := ProcessGame(i, 10, enums.Large, 99)

        assert.Equal(t, int16(i), result.CardDrawn)
        assert.Equal(t, enums.W, result.GameResult)
        assert.Equal(t, int64(20), result.MoneyInPlay)
        assert.Equal(t, int64(99), result.RemainingBalance)
    }
}
