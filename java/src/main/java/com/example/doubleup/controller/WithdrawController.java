package com.example.doubleup.controller;

import com.example.doubleup.dto.WithdrawRequestDto;
import com.example.doubleup.model.Player;
import com.example.doubleup.service.AccessLogService;
import com.example.doubleup.service.PlayService;
import com.example.doubleup.service.WithdrawService;
import com.example.doubleup.validation.WithdrawValidator;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Optional;

@RestController
@RequestMapping("/withdraw")
public class WithdrawController {

    @Autowired
    private PlayService playService;

    @Autowired
    private WithdrawService withdrawService;

    @Autowired
    private AccessLogService accessLogService;

    @PostMapping("/withdrawmoney")
    public String withdraw(HttpServletRequest request,  @RequestBody WithdrawRequestDto requestBody) throws Exception {
        accessLogService.writeAccessLog(request.getRemoteAddr(), "/withdraw/withdrawmoney");

        // player should exist in DB
        Player player = playService.getPlayer(requestBody.getPlayerId());

        // validate request dto
        WithdrawValidator.validateWithdrawRequest(requestBody, player);

        return withdrawService.withdraw(player);
    }
}
