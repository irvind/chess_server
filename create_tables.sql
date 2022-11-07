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
    created_by bigint NOT NULL,
    opponent bigint,
    creator_white boolean,
    created_at timestamp NOT NULL
);
