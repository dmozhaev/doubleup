package validation

import (
    "errors"
    "strings"
    "double_up/model"
    "double_up/dto"
)

func ValidateWithdrawRequest(requestDto dto.WithdrawRequestDto, player *model.Player) error {
	var dtoErrors []string

	// money should be in play already
	if player.MoneyInPlay == 0 {
		dtoErrors = append(dtoErrors, "WithdrawValidator: money should be in play already!")
	}

	if len(dtoErrors) > 0 {
		return errors.New(strings.Join(dtoErrors, ", "))
	}

	return nil
}
