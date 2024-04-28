package com.example.doubleup.dto;

import com.example.doubleup.enums.SmallLargeChoice;

import java.util.UUID;

public class PlayContinueRequestDto {
    private UUID playerId;

    private SmallLargeChoice choice;

    public PlayContinueRequestDto() {}

    public UUID getPlayerId() {
        return playerId;
    }

    public void setPlayerId(UUID playerId) {
        this.playerId = playerId;
    }

    public SmallLargeChoice getChoice() {
        return choice;
    }

    public void setChoice(SmallLargeChoice choice) {
        this.choice = choice;
    }
}
