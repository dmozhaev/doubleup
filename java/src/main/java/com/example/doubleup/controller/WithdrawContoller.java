package com.example.doubleup.controller;

import com.example.doubleup.dao.PlayerRepository;
import com.example.doubleup.dto.WithdrawRequestDto;
import com.example.doubleup.model.Player;
import com.example.doubleup.service.WithdrawService;
import com.example.doubleup.validation.WithdrawValidator;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Optional;

@RestController
@RequestMapping("/withdraw")
public class WithdrawContoller {

    @Autowired
    private PlayerRepository playerRepository;

    @Autowired
    private WithdrawService withdrawService;

    @PostMapping("/withdrawmoney")
    public String withdraw(@RequestBody WithdrawRequestDto requestBody) throws Exception {

        // player should exist in DB
        Optional<Player> playerOptional = playerRepository.findById(requestBody.getPlayerId());
        if (playerOptional.isEmpty()) {
            throw new Exception("Player id: " + requestBody.getPlayerId() + " not found!");
        }
        Player player = playerOptional.get();

        // validate request dto
        WithdrawValidator.validateWithdrawRequest(requestBody, player);

        return withdrawService.withdraw(player);
    }
}
