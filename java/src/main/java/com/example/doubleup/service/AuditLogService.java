package com.example.doubleup.service;

import com.example.doubleup.dao.AuditLogRepository;
import com.example.doubleup.enums.AuditOperation;
import com.example.doubleup.model.AuditLog;
import com.example.doubleup.model.Player;
import jakarta.transaction.Transactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.UUID;

@Service
public class AuditLogService {

    @Autowired
    private AuditLogRepository auditLogRepository;

    @Transactional
    public void writeAuditLog(Player player, AuditOperation operation, UUID recordId, String targetTable) {
        AuditLog auditLog = new AuditLog(player, recordId, targetTable, LocalDateTime.now(), operation);
        auditLogRepository.save(auditLog);
    }
}
