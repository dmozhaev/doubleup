-- create db and connect to it
CREATE DATABASE doubleupdb;
\c doubleupdb;

-- extension for generating uuid:s
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- create a database user
CREATE USER doubleup WITH PASSWORD 'doubleup';

-- create tables and relations
CREATE TABLE player (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    money_in_play BIGINT NOT NULL,  -- money is stored in cents
	account_balance BIGINT NOT NULL,
	CONSTRAINT chk_money_in_play_non_negative CHECK (money_in_play >= 0),  -- allowing play money -feature
	CONSTRAINT chk_account_balance_non_negative CHECK (account_balance >= 0)
);

CREATE TABLE game (
    id UUID PRIMARY KEY,
	player_id UUID NOT NULL,
	created_at TIMESTAMP DEFAULT now(),
	bet_size BIGINT NOT NULL,
	player_choice VARCHAR(100) NOT NULL,
	card_drawn SMALLINT NOT NULL,
	potential_profit BIGINT NOT NULL,
	FOREIGN KEY (player_id) REFERENCES player (id),
	CONSTRAINT chk_bet_size_non_negative CHECK (bet_size >= 0),
	CONSTRAINT chk_card_drawn_between CHECK (card_drawn BETWEEN 1 AND 13),
	CONSTRAINT chk_potential_profit_non_negative CHECK (potential_profit >= 0)
);

CREATE TABLE withdrawal (
    id UUID PRIMARY KEY,
	player_id UUID NOT NULL,
	created_at TIMESTAMP DEFAULT now(),
	amount BIGINT NOT NULL,
	FOREIGN KEY (player_id) REFERENCES player (id),
	CONSTRAINT chk_amount_positive CHECK (amount > 0)
);

-- indexes for potentially frequently used queries 
CREATE INDEX idx_player_name ON player (name);
CREATE INDEX idx_game_created_at ON game (created_at);
CREATE INDEX idx_game_bet_size ON game (bet_size);
CREATE INDEX idx_withdrawal_created_at ON withdrawal (created_at);
CREATE INDEX idx_withdrawal_amount ON withdrawal (amount);

-- insert test data (seed)
INSERT INTO player (id, name, money_in_play, account_balance) VALUES
(uuid_generate_v4(), 'heikki', 0, 1000),
(uuid_generate_v4(), 'gambler', 0, 100);

-- grant privileges to the user
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE player TO doubleup;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE game TO doubleup;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE withdrawal TO doubleup;
