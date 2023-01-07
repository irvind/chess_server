DROP TABLE IF EXISTS moves;
DROP TABLE IF EXISTS games;
DROP TABLE IF EXISTS players;

CREATE TABLE players (
    id bigserial PRIMARY KEY,
    name varchar(64) NOT NULL,
    auth_secret varchar(64) NOT NULL,
    created_at timestamp NOT NULL
);

CREATE TABLE games (
    id bigserial PRIMARY KEY,
    created_by bigint NOT NULL REFERENCES players,
    opponent bigint,
    creator_white boolean,
    created_at timestamp NOT NULL,
    finished boolean DEFAULT false
);

CREATE TABLE moves (
    id bigserial PRIMARY KEY,
    game_id bigint NOT NULL REFERENCES games,
    idx int NOT NULL,
    description varchar(8) NOT NULL,
    created_at timestamp NOT NULL
);
