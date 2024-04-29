package com.example.doubleup.controller;

import com.example.doubleup.dto.PlayContinueRequestDto;
import com.example.doubleup.dto.PlayResponseDto;
import com.example.doubleup.dto.PlayStartRequestDto;
import com.example.doubleup.model.Player;
import com.example.doubleup.service.AccessLogService;
import com.example.doubleup.service.PlayService;
import com.example.doubleup.validation.PlayValidator;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/play")
public class PlayController {

    @Autowired
    private PlayService playService;

    @Autowired
    private AccessLogService accessLogService;

    @PostMapping("/start")
    public PlayResponseDto playStart(HttpServletRequest request, @RequestBody PlayStartRequestDto requestBody) throws Exception {
        accessLogService.checkAccessAllowed(request.getRemoteAddr(), "/play/start");

        // player should exist in DB
        Player player = playService.getPlayer(requestBody.getPlayerId());

        // validate request dto
        PlayValidator.validateStartRequest(requestBody, player);

        return playService.startGame(player, requestBody.getBetSize(), requestBody.getChoice());
    }

    @PostMapping("/continue")
    public PlayResponseDto playContinue(HttpServletRequest request, @RequestBody PlayContinueRequestDto requestBody) throws Exception {
        accessLogService.checkAccessAllowed(request.getRemoteAddr(), "/play/continue");

        // player should exist in DB
        Player player = playService.getPlayer(requestBody.getPlayerId());

        // validate request dto
        PlayValidator.validateContinueRequest(requestBody, player);

        return playService.continueGame(player, player.getMoneyInPlay(), requestBody.getChoice());
    }
}
