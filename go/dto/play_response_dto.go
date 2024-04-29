package dto

type PlayResponseDto struct {
    CardDrawn       int16
    GameResult      GameResult
    MoneyInPlay     int64
    RemainingBalance int64
}

func NewPlayResponseDto(cardDrawn int16, gameResult GameResult, moneyInPlay, remainingBalance int64) *PlayResponseDto {
    return &PlayResponseDto{
        CardDrawn:       cardDrawn,
        GameResult:      gameResult,
        MoneyInPlay:     moneyInPlay,
        RemainingBalance: remainingBalance,
    }
}
