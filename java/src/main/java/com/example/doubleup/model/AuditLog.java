package com.example.doubleup.model;

import com.example.doubleup.enums.AuditOperation;
import jakarta.persistence.Entity;
import jakarta.persistence.Enumerated;
import jakarta.persistence.EnumType;
import jakarta.persistence.FetchType;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.Id;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.Table;
import org.hibernate.annotations.GenericGenerator;

import java.time.LocalDateTime;
import java.util.UUID;

@Entity
@Table(name = "audit_log")
public class AuditLog {
    @Id
    @GeneratedValue(generator = "uuid2")
    @GenericGenerator(name = "uuid2", strategy = "org.hibernate.id.UUIDGenerator")
    private UUID id;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "player_id")
    private Player player;

    private UUID recordId;

    private LocalDateTime createdAt;

    @Enumerated(EnumType.STRING)
    private AuditOperation operation;

    public AuditLog() {
    }

    public AuditLog(Player player, UUID recordId, LocalDateTime createdAt, AuditOperation operation) {
        this.player = player;
        this.recordId = recordId;
        this.createdAt = createdAt;
        this.operation = operation;
    }

    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }

    public Player getPlayer() {
        return player;
    }

    public void setPlayer(Player player) {
        this.player = player;
    }

    public UUID getRecordId() {
        return recordId;
    }

    public void setRecordId(UUID recordId) {
        this.recordId = recordId;
    }

    public LocalDateTime getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(LocalDateTime createdAt) {
        this.createdAt = createdAt;
    }

    public AuditOperation getOperation() {
        return operation;
    }

    public void setOperation(AuditOperation operation) {
        this.operation = operation;
    }
}
