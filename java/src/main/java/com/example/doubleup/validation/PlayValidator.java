package com.example.doubleup.validation;

import com.example.doubleup.dto.PlayContinueRequestDto;
import com.example.doubleup.dto.PlayStartRequestDto;
import com.example.doubleup.model.Player;

import java.util.ArrayList;
import java.util.List;

public class PlayValidator {
    public static void validateStartRequest(PlayStartRequestDto requestBody, Player player) throws Exception {
        List<String> errors = new ArrayList<>();

        // there should be no money in play in order to start
        if (player.getMoneyInPlay() != 0) {
            errors.add("PlayValidator: there should be no money in play in order to start!");
        }

        // validate bet size
        if (requestBody.getBetSize() <= 0) {
            errors.add("PlayValidator: bet is too small");
        }
        if (requestBody.getBetSize() > player.getAccountBalance()) {
            errors.add("PlayValidator: bet is too large, insufficient funds");;
        }

        if (!errors.isEmpty()) {
            throw new Exception(String.join(", ", errors));
        }
    }

    public static void validateContinueRequest(PlayContinueRequestDto requestBody, Player player) throws Exception {
        List<String> errors = new ArrayList<>();

        // money should be in play already
        if (player.getMoneyInPlay() == 0) {
            errors.add("PlayValidator: money should be in play already!");
        }

        if (!errors.isEmpty()) {
            throw new Exception(String.join(", ", errors));
        }
    }
}
