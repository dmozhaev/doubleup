package com.example.doubleup.controller;

import com.example.doubleup.dao.PlayerRepository;
import com.example.doubleup.dto.PlayContinueRequestDto;
import com.example.doubleup.dto.PlayResponseDto;
import com.example.doubleup.dto.PlayStartRequestDto;
import com.example.doubleup.enums.GameResult;
import com.example.doubleup.model.Player;
import com.example.doubleup.service.PlayService;
import com.example.doubleup.validation.PlayValidator;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Optional;

@RestController
@RequestMapping("/play")
public class PlayController {

    @Autowired
    private PlayerRepository playerRepository;

    @Autowired
    private PlayService playService;

    @PostMapping("/start")
    public PlayResponseDto playStart(@RequestBody PlayStartRequestDto requestBody) throws Exception {

        // player should exist in DB
        Optional<Player> playerOptional = playerRepository.findById(requestBody.getPlayerId());
        if (playerOptional.isEmpty()) {
            throw new Exception("Player id: " + requestBody.getPlayerId() + " not found!");
        }
        Player player = playerOptional.get();

        // validate request dto
        PlayValidator.validateStartRequest(requestBody, player);

        return playService.startGame(player, requestBody.getBetSize(), requestBody.getChoice());
    }

    @PostMapping("/continue")
    public PlayResponseDto playContinue(@RequestBody PlayContinueRequestDto requestBody) throws Exception {

        // player should exist in DB
        Optional<Player> playerOptional = playerRepository.findById(requestBody.getPlayerId());
        if (playerOptional.isEmpty()) {
            throw new Exception("Player id: " + requestBody.getPlayerId() + " not found!");
        }
        Player player = playerOptional.get();

        // validate request dto
        PlayValidator.validateContinueRequest(requestBody, player);

        return playService.continueGame(player, player.getMoneyInPlay(), requestBody.getChoice());
    }
}
