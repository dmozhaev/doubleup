package com.example.doubleup.model;

import com.example.doubleup.enums.GameResult;
import com.example.doubleup.enums.SmallLargeChoice;
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

import java.time.OffsetDateTime;
import java.util.UUID;

@Entity
@Table(name = "game")
public class Game {
    @Id
    @GeneratedValue(generator = "uuid2")
    @GenericGenerator(name = "uuid2", strategy = "org.hibernate.id.UUIDGenerator")
    private UUID id;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "player_id")
    private Player player;

    private OffsetDateTime createdAt;

    private Long betSize;

    @Enumerated(EnumType.STRING)
    private SmallLargeChoice playerChoice;

    private Short cardDrawn;

    private Long potentialProfit;

    @Enumerated(EnumType.STRING)
    private GameResult gameResult;

    public Game() {
    }

    public Game(Player player, OffsetDateTime createdAt, Long betSize, SmallLargeChoice playerChoice, Short cardDrawn, Long potentialProfit, GameResult gameResult) {
        this.player = player;
        this.createdAt = createdAt;
        this.betSize = betSize;
        this.playerChoice = playerChoice;
        this.cardDrawn = cardDrawn;
        this.potentialProfit = potentialProfit;
        this.gameResult = gameResult;
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

    public OffsetDateTime getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(OffsetDateTime createdAt) {
        this.createdAt = createdAt;
    }

    public Long getBetSize() {
        return betSize;
    }

    public void setBetSize(Long betSize) {
        this.betSize = betSize;
    }

    public SmallLargeChoice getPlayerChoice() {
        return playerChoice;
    }

    public void setPlayerChoice(SmallLargeChoice playerChoice) {
        this.playerChoice = playerChoice;
    }

    public Short getCardDrawn() {
        return cardDrawn;
    }

    public void setCardDrawn(Short cardDrawn) {
        this.cardDrawn = cardDrawn;
    }

    public Long getPotentialProfit() {
        return potentialProfit;
    }

    public void setPotentialProfit(Long potentialProfit) {
        this.potentialProfit = potentialProfit;
    }

    public GameResult getGameResult() {
        return gameResult;
    }

    public void setGameResult(GameResult gameResult) {
        this.gameResult = gameResult;
    }
}
