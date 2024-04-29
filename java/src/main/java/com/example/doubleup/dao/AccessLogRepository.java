package com.example.doubleup.dao;

import com.example.doubleup.model.AccessLog;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.time.OffsetDateTime;
import java.util.UUID;

@Repository
public interface AccessLogRepository extends JpaRepository<AccessLog, UUID> {

    @Query("SELECT COUNT(e) FROM AccessLog e WHERE e.createdAt >= :startTime and ipAddress = :ipAddress")
    int countRowsForLastMinute(String ipAddress, OffsetDateTime startTime);
}
