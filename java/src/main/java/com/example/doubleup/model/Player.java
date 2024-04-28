package com.example.doubleup.model;

import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

import java.util.UUID;

@Entity
@Table(name = "player")
public class Player {
    @Id
    private UUID id;

    private String name;

    private Long moneyInPlay;

    private Long accountBalance;

    public Player() {
    }

    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Long getMoneyInPlay() {
        return moneyInPlay;
    }

    public void setMoneyInPlay(Long moneyInPlay) {
        this.moneyInPlay = moneyInPlay;
    }

    public Long getAccountBalance() {
        return accountBalance;
    }

    public void setAccountBalance(Long accountBalance) {
        this.accountBalance = accountBalance;
    }
}
