package validation

import (
    "errors"
    "strings"
    "double_up/model"
    "double_up/dto"
)

func ValidateStartRequest(requestDto dto.PlayStartRequestDto, player *model.Player) error {
	var dtoErrors []string

	// There should be no money in play in order to start
	if player.MoneyInPlay != 0 {
		dtoErrors = append(dtoErrors, "PlayValidator: there should be no money in play in order to start!")
	}

	// Validate bet size
	if requestDto.BetSize <= 0 {
		dtoErrors = append(dtoErrors, "PlayValidator: bet is too small")
	}
	if requestDto.BetSize > player.AccountBalance {
		dtoErrors = append(dtoErrors, "PlayValidator: bet is too large, insufficient funds")
	}

	if len(dtoErrors) > 0 {
		return errors.New(strings.Join(dtoErrors, ", "))
	}

	return nil
}
