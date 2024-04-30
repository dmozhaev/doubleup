package service

import (
    "database/sql"
    "github.com/google/uuid"
    "math/rand"
    "time"
    "double_up/dao"
    "double_up/dto"
    "double_up/enums"
    "double_up/model"
)

func GetPlayer(db *sql.DB, id uuid.UUID) (*model.Player, error) {
    player, err := dao.FindPlayerById(db, id)
    WriteAuditLog(db, player, enums.Select, player.ID, "player")
    return player, err
}

func ProcessGame(db *sql.DB, betSize int64, playerChoice enums.SmallLargeChoice, accountBalance int64) *dto.PlayResponseDto {
    rand.Seed(time.Now().UnixNano())
    randomNumber := rand.Intn(13) + 1

    // generate a random number to ensure cryptographic-strength randomness
	var gameChoice enums.SmallLargeChoice
	if randomNumber <= 6 {
		gameChoice = enums.Small
	} else if randomNumber >= 8 {
		gameChoice = enums.Large
	}

    // decide the game result
	gameResult := enums.L
	if gameChoice == playerChoice {
		gameResult = enums.W
	}

    // money in play
	var moneyInPlay int64
	if gameResult == enums.W {
		moneyInPlay = betSize * 2
	} else {
		moneyInPlay = 0
	}

	return &dto.PlayResponseDto{
		CardDrawn:      int16(randomNumber),
		GameResult:     gameResult,
		MoneyInPlay:    moneyInPlay,
		RemainingBalance: accountBalance,
	}
}

func PlayGame(db *sql.DB, player *model.Player, betSize int64, choice enums.SmallLargeChoice) (*dto.PlayResponseDto, error) {
    playResponseDto := ProcessGame(db, betSize, choice, player.AccountBalance)
    player.MoneyInPlay = playResponseDto.MoneyInPlay
    dao.UpdatePlayer(db, player)
    WriteAuditLog(db, player, enums.Update, player.ID, "player")

    //WriteAuditLog(db, player, enums.Insert, game.ID, "game")

    return playResponseDto, nil
}

func StartGame(db *sql.DB, player *model.Player, betSize int64, choice enums.SmallLargeChoice) (*dto.PlayResponseDto, error) {
    player.AccountBalance -= betSize
    return PlayGame(db, player, betSize, choice)
}

func ContinueGame(db *sql.DB, player *model.Player, betSize int64, choice enums.SmallLargeChoice) (*dto.PlayResponseDto, error) {
    return PlayGame(db, player, betSize, choice)
}