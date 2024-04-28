package com.example.doubleup.dto;

public class PlayStartRequestDto extends PlayContinueRequestDto {
    private Long betSize;

    public PlayStartRequestDto() {}

    public Long getBetSize() {
        return betSize;
    }

    public void setBetSize(Long betSize) {
        this.betSize = betSize;
    }
}
