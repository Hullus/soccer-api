-- Data Types
CREATE TYPE listing_status AS ENUM ('active', 'sold', 'cancelled');
CREATE TYPE player_position AS ENUM ('goalkeeper', 'defender', 'midfielder', 'attacker');

ALTER type listing_status owner TO postgres;
ALTER type player_position owner TO postgres;

-- Tables
CREATE table users
(
    id            BIGSERIAL PRIMARY KEY,
    email         varchar(255)                          NOT NULL UNIQUE,
    password_hash varchar(255)                          NOT NULL,
    created_at    TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);
ALTER table users
    owner TO postgres;

CREATE table teams
(
    id           BIGSERIAL PRIMARY KEY,
    name         varchar(100)                   NOT NULL,
    budget_cents VARCHAR     default 5000000.00 NOT NULL,
    created_at   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    country      VARCHAR(100)                   NOT NULL,
    owner_id     BIGINT
        constraint unique_owner_id UNIQUE
        constraint owner_id references users ON DELETE SET NULL
);
ALTER table teams
    owner TO postgres;

CREATE table players
(
    id                 BIGSERIAL
        CONSTRAINT player_pkey PRIMARY KEY,
    team_id            BIGINT
        constraint team_id references teams ON DELETE
            SET
            NULL,
    first_name         varchar(100)             NOT NULL,
    last_name          varchar(100)             NOT NULL,
    country            varchar(100)             NOT NULL,
    age                integer                  NOT NULL
        constraint player_age_check check ((age >= 18) AND (age <= 40)),
    market_value_cents BIGINT      default 0.00 NOT NULL,
    created_at         TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    position           player_position          NOT NULL
);
ALTER table players
    owner TO postgres;

CREATE table transfer_listings
(
    id                 BIGSERIAL
        constraint table_name_pkey PRIMARY KEY,
    player_id          BIGINT                                           NOT NULL
        constraint table_name_player_id_key UNIQUE
        constraint table_name_player_id_fkey references players,
    seller_team_id     BIGINT                                           NOT NULL REFERENCES teams (id),
    sold_to_team_id    BIGINT REFERENCES teams (id),
    status             listing_status default 'active':: listing_status NOT NULL,
    asking_price_cents BIGINT                                           NOT NULL
        constraint table_name_asking_price_check check ( asking_price_cents > (0):: numeric ),
    listed_at          TIMESTAMPTZ    DEFAULT CURRENT_TIMESTAMP         NOT NULL,
    sold_at            TIMESTAMPTZ    DEFAULT CURRENT_TIMESTAMP
);
ALTER table transfer_listings
    owner TO postgres;