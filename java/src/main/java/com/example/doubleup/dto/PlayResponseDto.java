package com.example.doubleup.dto;

import com.example.doubleup.enums.GameResult;

public class PlayResponseDto {
    private Short cardDrawn;

    private GameResult gameResult;

    private Long moneyInPlay;

    private Long remainingBalance;

    public PlayResponseDto(Short cardDrawn, GameResult gameResult, Long moneyInPlay, Long remainingBalance) {
        this.cardDrawn = cardDrawn;
        this.gameResult = gameResult;
        this.moneyInPlay = moneyInPlay;
        this.remainingBalance = remainingBalance;
    }

    public Short getCardDrawn() {
        return cardDrawn;
    }

    public void setCardDrawn(Short cardDrawn) {
        this.cardDrawn = cardDrawn;
    }

    public GameResult getGameResult() {
        return gameResult;
    }

    public void setGameResult(GameResult gameResult) {
        this.gameResult = gameResult;
    }

    public Long getMoneyInPlay() {
        return moneyInPlay;
    }

    public void setMoneyInPlay(Long moneyInPlay) {
        this.moneyInPlay = moneyInPlay;
    }

    public Long getRemainingBalance() {
        return remainingBalance;
    }

    public void setRemainingBalance(Long remainingBalance) {
        this.remainingBalance = remainingBalance;
    }
}
