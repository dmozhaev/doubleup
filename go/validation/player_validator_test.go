package validation

import (
	"double_up/enums"
	"double_up/model"
	"double_up/dto"
	"strings"
	"testing"
)

func TestValidateStartRequest(t *testing.T) {
	// Test case 1: No money in play, valid bet size
	player1 := &model.Player{MoneyInPlay: 0, AccountBalance: 100}
	requestDto1 := dto.PlayStartRequestDto{BetSize: 50, Choice: enums.Large}
	if err := ValidateStartRequest(requestDto1, player1); err != nil {
		t.Errorf("Test case 1 failed: Expected no error, got %v", err)
	}

	// Test case 2: Money in play, valid bet size
	player2 := &model.Player{MoneyInPlay: 50, AccountBalance: 100}
	requestDto2 := dto.PlayStartRequestDto{BetSize: 50, Choice: enums.Small}
	expectedErr2 := "PlayValidator: there should be no money in play in order to start!"
	if err := ValidateStartRequest(requestDto2, player2); err == nil || err.Error() != expectedErr2 {
		t.Errorf("Test case 2 failed: Expected error '%s', got %v", expectedErr2, err)
	}

	// Test case 3: No money in play, valid bet size exceeds account balance
	player3 := &model.Player{MoneyInPlay: 0, AccountBalance: 100}
	requestDto3 := dto.PlayStartRequestDto{BetSize: 150, Choice: enums.Large}
	expectedErr3 := "PlayValidator: bet is too large, insufficient funds"
	if err := ValidateStartRequest(requestDto3, player3); err == nil || err.Error() != expectedErr3 {
		t.Errorf("Test case 3 failed: Expected error '%s', got %v", expectedErr3, err)
	}

	// Test case 4: Invalid bet size (zero)
	player4 := &model.Player{MoneyInPlay: 0, AccountBalance: 100}
	requestDto4 := dto.PlayStartRequestDto{BetSize: 0, Choice: enums.Small}
	expectedErr4 := "PlayValidator: bet is too small"
	if err := ValidateStartRequest(requestDto4, player4); err == nil || err.Error() != expectedErr4 {
		t.Errorf("Test case 4 failed: Expected error '%s', got %v", expectedErr4, err)
	}

	// Test case 5: Invalid bet size (negative value)
	player5 := &model.Player{MoneyInPlay: 0, AccountBalance: 100}
	requestDto5 := dto.PlayStartRequestDto{BetSize: -1, Choice: enums.Large}
	expectedErr5 := "PlayValidator: bet is too small"
	if err := ValidateStartRequest(requestDto5, player5); err == nil || err.Error() != expectedErr4 {
		t.Errorf("Test case 5 failed: Expected error '%s', got %v", expectedErr5, err)
	}

	// Test case 6: Money in play, valid bet size exceeds account balance
	player6 := &model.Player{MoneyInPlay: 50, AccountBalance: 100}
	requestDto6 := dto.PlayStartRequestDto{BetSize: 150, Choice: enums.Small}
	expectedErr6 := "PlayValidator: there should be no money in play in order to start!, PlayValidator: bet is too large, insufficient funds"
	if err := ValidateStartRequest(requestDto6, player6); err == nil || !strings.Contains(err.Error(), expectedErr6) {
		t.Errorf("Test case 6 failed: Expected error containing '%s', got %v", expectedErr6, err)
	}

	// Test case 7: Invalid player choice
	player7 := &model.Player{MoneyInPlay: 0, AccountBalance: 100}
	requestDto7 := dto.PlayStartRequestDto{BetSize: 10}
	expectedErr7 := "PlayValidator: choice is invalid"
	if err := ValidateStartRequest(requestDto7, player7); err == nil || !strings.Contains(err.Error(), expectedErr7) {
		t.Errorf("Test case 7 failed: Expected error containing '%s', got %v", expectedErr7, err)
	}
}

func TestValidateContinueRequest(t *testing.T) {
	// Test case 1: Money in play, valid choice
	player1 := &model.Player{MoneyInPlay: 50}
	requestDto1 := dto.PlayContinueRequestDto{Choice: enums.Small}
	if err := ValidateContinueRequest(requestDto1, player1); err != nil {
		t.Errorf("Test case 1 failed: Expected no error, got %v", err)
	}

	// Test case 2: No money in play
	player2 := &model.Player{MoneyInPlay: 0}
	requestDto2 := dto.PlayContinueRequestDto{Choice: enums.Large}
	expectedErr2 := "PlayValidator: money should be in play already!"
	if err := ValidateContinueRequest(requestDto2, player2); err == nil || err.Error() != expectedErr2 {
		t.Errorf("Test case 2 failed: Expected error '%s', got %v", expectedErr2, err)
	}

	// Test case 3: Invalid player choice
	player3 := &model.Player{MoneyInPlay: 50}
	requestDto3 := dto.PlayContinueRequestDto{}
	expectedErr3 := "PlayValidator: choice is invalid"
	if err := ValidateContinueRequest(requestDto3, player3); err == nil || err.Error() != expectedErr3 {
		t.Errorf("Test case 3 failed: Expected error '%s', got %v", expectedErr3, err)
	}
}
