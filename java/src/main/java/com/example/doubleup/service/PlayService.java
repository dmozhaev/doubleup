package com.example.doubleup.service;

import com.example.doubleup.dao.GameRepository;
import com.example.doubleup.dao.PlayerRepository;
import com.example.doubleup.dto.PlayResponseDto;
import com.example.doubleup.enums.GameResult;
import com.example.doubleup.enums.SmallLargeChoice;
import com.example.doubleup.model.Game;
import com.example.doubleup.model.Player;
import jakarta.transaction.Transactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.security.SecureRandom;
import java.time.LocalDateTime;
import java.util.Optional;
import java.util.UUID;

@Service
public class PlayService {

    @Autowired
    private PlayerRepository playerRepository;

    @Autowired
    private GameRepository gameRepository;

    public Player getPlayer(UUID id) throws Exception {
        Optional<Player> playerOptional = playerRepository.findById(id);
        if (playerOptional.isEmpty()) {
            throw new Exception("Player id: " + id + " not found!");
        }
        return playerOptional.get();
    }

    private PlayResponseDto processGame(Long betSize, SmallLargeChoice playerChoice, Long accountBalance) {
        SecureRandom secureRandom = new SecureRandom();
        int randomNumber = secureRandom.nextInt(13) + 1;

        // generate a random number to ensure cryptographic-strength randomness
        SmallLargeChoice gameChoice = null;
        if (randomNumber <= 6) {
            gameChoice = SmallLargeChoice.SMALL;
        } else if (randomNumber >= 8) {
            gameChoice = SmallLargeChoice.LARGE;
        }

        // decide the game result
        GameResult gameResult = GameResult.L;
        if (gameChoice == playerChoice) {
            gameResult = GameResult.W;
        }

        // money in play
        Long moneyInPlay = gameResult == GameResult.W ? betSize * 2 : 0;

        return new PlayResponseDto((short)randomNumber, gameResult, moneyInPlay, accountBalance);
    }

    public PlayResponseDto playGame(Player player, Long betSize, SmallLargeChoice choice) {
        PlayResponseDto playResponseDto = processGame(betSize, choice, player.getAccountBalance());
        player.setMoneyInPlay(playResponseDto.getMoneyInPlay());
        playerRepository.save(player);
        Game game = new Game(player, LocalDateTime.now(), betSize, choice, playResponseDto.getCardDrawn(), betSize * 2, playResponseDto.getGameResult());
        gameRepository.save(game);
        return playResponseDto;
    }

    @Transactional
    public PlayResponseDto startGame(Player player, Long betSize, SmallLargeChoice choice) {
        player.setAccountBalance(player.getAccountBalance() - betSize);
        return playGame(player, betSize, choice);
    }

    @Transactional
    public PlayResponseDto continueGame(Player player, Long betSize, SmallLargeChoice choice) {
        return playGame(player, betSize, choice);
    }
}
