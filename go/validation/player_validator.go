package validation

import (
    "errors"
    "strings"
    "double_up/model"
    "double_up/dto"
    "double_up/enums"
)

func ValidateStartRequest(requestDto dto.PlayStartRequestDto, player *model.Player) error {
	var dtoErrors []string

	// there should be no money in play in order to start
	if player.MoneyInPlay != 0 {
		dtoErrors = append(dtoErrors, "PlayValidator: there should be no money in play in order to start!")
	}

	// validate bet size
	if requestDto.BetSize <= 0 {
		dtoErrors = append(dtoErrors, "PlayValidator: bet is too small")
	}
	if requestDto.BetSize > player.AccountBalance {
		dtoErrors = append(dtoErrors, "PlayValidator: bet is too large, insufficient funds")
	}

	// invalid choice
	if requestDto.Choice != enums.Small && requestDto.Choice != enums.Large {
		dtoErrors = append(dtoErrors, "PlayValidator: choice is invalid")
	}

	if len(dtoErrors) > 0 {
		return errors.New(strings.Join(dtoErrors, ", "))
	}

	return nil
}

func ValidateContinueRequest(requestDto dto.PlayContinueRequestDto, player *model.Player) error {
	var dtoErrors []string

	// money should be in play already
	if player.MoneyInPlay == 0 {
		dtoErrors = append(dtoErrors, "PlayValidator: money should be in play already!")
	}

	// invalid choice
	if requestDto.Choice != enums.Small && requestDto.Choice != enums.Large {
		dtoErrors = append(dtoErrors, "PlayValidator: choice is invalid")
	}

	if len(dtoErrors) > 0 {
		return errors.New(strings.Join(dtoErrors, ", "))
	}

	return nil
}
