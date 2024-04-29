package com.example.doubleup.service;

import com.example.doubleup.Constants;
import com.example.doubleup.dao.AccessLogRepository;
import com.example.doubleup.model.AccessLog;
import jakarta.transaction.Transactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.OffsetDateTime;
import java.time.temporal.ChronoUnit;

@Service
public class AccessLogService {

    @Autowired
    private AccessLogRepository accessLogRepository;

    @Transactional
    private void writeAccessLog(String ipAddress, String api) {
        AccessLog accessLog = new AccessLog(OffsetDateTime.now(), ipAddress, api);
        accessLogRepository.save(accessLog);
    }

    private void checkAccess(String ipAddress) throws Exception {
        OffsetDateTime startTime = OffsetDateTime.now().minus(1, ChronoUnit.MINUTES);
        int apiCountLastMinute = accessLogRepository.countRowsForLastMinute(ipAddress, startTime);
        if (apiCountLastMinute > Constants.REQUEST_LIMIT_PER_MINUTE) {
            throw new Exception("Too many requests! IP address: " + ipAddress);
        };
    }

    public void checkAccessAllowed(String ipAddress, String api) throws Exception {
        this.writeAccessLog(ipAddress, api);
        this.checkAccess(ipAddress);
    }
}
