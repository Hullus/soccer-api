INSERT INTO users (email, password_hash)
VALUES ('manager1@test.com', '$2a$10$ByI67ZghS099v+F9V.BDe.jQkU6Qf7XvMvU.yK7H7W1M7M7M7M7M7'),
       ('manager2@test.com', '$2a$10$ByI67ZghS099v+F9V.BDe.jQkU6Qf7XvMvU.yK7H7W1M7M7M7M7M7');

INSERT INTO teams (name, country, owner_id)
VALUES ('FC Barcelona', 'Spain', 1),
       ('Manchester City', 'England', 2);

INSERT INTO players (team_id, first_name, last_name, country, age, position)
VALUES (1, 'Marc-André', 'ter Stegen', 'Germany', 31, 'goalkeeper'),
       (1, 'Ronald', 'Araújo', 'Uruguay', 24, 'defender'),
       (1, 'Gavi', 'Páez', 'Spain', 19, 'midfielder'),
       (1, 'Robert', 'Lewandowski', 'Poland', 35, 'attacker');

INSERT INTO players (team_id, first_name, last_name, country, age, position)
VALUES (2, 'Ederson', 'Moraes', 'Brazil', 30, 'goalkeeper'),
       (2, 'Ruben', 'Dias', 'Portugal', 26, 'defender'),
       (2, 'Kevin', 'De Bruyne', 'Belgium', 32, 'midfielder'),
       (2, 'Erling', 'Haaland', 'Norway', 23, 'attacker');

INSERT INTO transfer_listings (player_id, seller_team_id, status, asking_price_cents)
VALUES (4, 1, 'active', 120000000); -- Listing Lewandowski for 1.2M