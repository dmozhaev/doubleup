package com.example.doubleup.dao;

import com.example.doubleup.model.AccessLog;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.UUID;

@Repository
public interface AccessLogRepository extends JpaRepository<AccessLog, UUID> {
}
