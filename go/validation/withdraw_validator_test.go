package validation

import (
    "testing"
    "double_up/model"
    "double_up/dto"
)

func TestValidateWithdrawRequest(t *testing.T) {
	// Test case 1: Valid request (money in play)
	player1 := &model.Player{MoneyInPlay: 100}
	requestDto1 := dto.WithdrawRequestDto{}
	if err := ValidateWithdrawRequest(requestDto1, player1); err != nil {
		t.Errorf("Test case 1 failed: Expected no error, got %v", err)
	}

	// Invalid request (no money in play)
	player2 := &model.Player{MoneyInPlay: 0}
	requestDto2 := dto.WithdrawRequestDto{}
	expectedErr2 := "WithdrawValidator: money should be in play already!"
	if err := ValidateWithdrawRequest(requestDto2, player2); err == nil || err.Error() != expectedErr2 {
		t.Errorf("Test case 2 failed: Expected error '%s', got %v", expectedErr2, err)
	}
}
