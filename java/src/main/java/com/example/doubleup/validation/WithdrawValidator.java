package com.example.doubleup.validation;

import com.example.doubleup.dto.WithdrawRequestDto;
import com.example.doubleup.model.Player;

import java.util.ArrayList;
import java.util.List;

public class WithdrawValidator {
    public static void validateWithdrawRequest(WithdrawRequestDto requestBody, Player player) throws Exception {
        List<String> errors = new ArrayList<>();

        // money should be in play already
        if (player.getMoneyInPlay() == 0) {
            errors.add("WithdrawValidator: money should be in play already!");
        }

        if (!errors.isEmpty()) {
            throw new Exception(String.join(", ", errors));
        }
    }
}
