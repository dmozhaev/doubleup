package com.example.doubleup.dto;

import java.util.UUID;

public class WithdrawRequestDto {
    private UUID playerId;

    public WithdrawRequestDto() {}

    public UUID getPlayerId() {
        return playerId;
    }

    public void setPlayerId(UUID playerId) {
        this.playerId = playerId;
    }
}
