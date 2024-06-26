package com.example.doubleup.service;

import com.example.doubleup.dao.PlayerRepository;
import com.example.doubleup.dao.WithdrawalRepository;
import com.example.doubleup.enums.AuditOperation;
import com.example.doubleup.model.Player;
import com.example.doubleup.model.Withdrawal;
import jakarta.transaction.Transactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.OffsetDateTime;

@Service
public class WithdrawService {

    @Autowired
    private PlayerRepository playerRepository;

    @Autowired
    private WithdrawalRepository withdrawalRepository;

    @Autowired
    private AuditLogService auditLogService;

    @Transactional
    public String withdraw(Player player) {
        Withdrawal withdrawal = new Withdrawal(player, OffsetDateTime.now(), player.getMoneyInPlay());
        withdrawal = withdrawalRepository.save(withdrawal);
        auditLogService.writeAuditLog(player, AuditOperation.INSERT, withdrawal.getId(), "withdrawal");

        player.setAccountBalance(player.getAccountBalance() + player.getMoneyInPlay());
        player.setMoneyInPlay(0L);
        player = playerRepository.save(player);
        auditLogService.writeAuditLog(player, AuditOperation.UPDATE, player.getId(), "player");

        return "OK";
    }
}
