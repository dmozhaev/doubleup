package com.example.doubleup.service;

import com.example.doubleup.dao.AccessLogRepository;
import com.example.doubleup.model.AccessLog;
import jakarta.transaction.Transactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;

@Service
public class AccessLogService {

    @Autowired
    private AccessLogRepository accessLogRepository;

    @Transactional
    public void writeAccessLog(String ipAddress, String api) {
        AccessLog accessLog = new AccessLog(LocalDateTime.now(), ipAddress, api);
        accessLogRepository.save(accessLog);
    }
}
