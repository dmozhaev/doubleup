package com.example.doubleup.service;

import com.example.doubleup.dao.PlayerRepository;
import com.example.doubleup.dao.WithdrawalRepository;
import com.example.doubleup.model.Player;
import com.example.doubleup.model.Withdrawal;
import jakarta.transaction.Transactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;

@Service
public class WithdrawService {

    @Autowired
    private PlayerRepository playerRepository;

    @Autowired
    private WithdrawalRepository withdrawalRepository;

    @Transactional
    public String withdraw(Player player) {
        Withdrawal withdrawal = new Withdrawal(player, LocalDateTime.now(), player.getMoneyInPlay());
        withdrawalRepository.save(withdrawal);

        player.setAccountBalance(player.getAccountBalance() + player.getMoneyInPlay());
        player.setMoneyInPlay(0L);
        playerRepository.save(player);

        return "OK";
    }
}
