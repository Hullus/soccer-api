INSERT INTO users (email, password_hash)
VALUES ('manager1@test.com', '$2a$10$ByI67ZghS099v+F9V.BDe.jQkU6Qf7XvMvU.yK7H7W1M7M7M7M7M7'),
       ('manager2@test.com', '$2a$10$ByI67ZghS099v+F9V.BDe.jQkU6Qf7XvMvU.yK7H7W1M7M7M7M7M7'),
       ('manager3@test.com', '$2a$10$ByI67ZghS099v+F9V.BDe.jQkU6Qf7XvMvU.yK7H7W1M7M7M7M7M7');

INSERT INTO teams (name, country, owner_id)
VALUES ('FC Barcelona', 'Spain', 1),
       ('Manchester City', 'England', 2),
       ('Bayern Munich', 'Germany', 3);

INSERT INTO players (team_id, first_name, last_name, country, age, position, market_value_cents)
SELECT (floor(random() * 3) + 1)::int,
       (ARRAY ['Marc', 'Robert', 'Kevin', 'Erling', 'Manuel', 'Joshua', 'Vinicius', 'Jude', 'Harry', 'Kylian'])[floor(random() * 10) + 1],
       (ARRAY ['Ter Stegen', 'Lewandowski', 'De Bruyne', 'Haaland', 'Neuer', 'Kimmich', 'Junior', 'Bellingham', 'Kane', 'Mbappe'])[floor(random() * 10) + 1],
       (ARRAY ['Germany','Georgia', 'Poland', 'Belgium', 'Norway', 'Brazil', 'England', 'France', 'Spain'])[floor(random() * 8) + 1],
       (floor(random() * (35 - 18 + 1) + 18))::int,
       (ARRAY ['goalkeeper', 'defender', 'midfielder', 'attacker']::player_position[])[floor(random() * 4) + 1],
       1000000
FROM generate_series(1, 45);

INSERT INTO players (team_id, first_name, last_name, country, age, position, market_value_cents)
SELECT NULL,
       (ARRAY ['James', 'John', 'Robert', 'Michael', 'William', 'David', 'Richard', 'Joseph', 'Thomas', 'Charles', 'Daniel', 'Matthew', 'Anthony', 'Donald', 'Mark', 'Paul', 'Steven', 'Andrew', 'Kenneth', 'Joshua'])[floor(random() * 20) + 1],
       (ARRAY ['Smith', 'Johnson', 'Williams', 'Jones', 'Brown', 'Davis', 'Miller', 'Wilson', 'Moore', 'Taylor', 'Anderson', 'Thomas', 'Jackson', 'White', 'Harris', 'Martin', 'Thompson', 'Garcia', 'Martinez', 'Robinson'])[floor(random() * 20) + 1],
       (ARRAY ['England', 'Spain', 'Germany', 'France', 'Italy', 'Brazil', 'Argentina', 'Netherlands', 'Portugal', 'Belgium', 'USA', 'Croatia', 'Uruguay', 'Colombia', 'Japan'])[floor(random() * 15) + 1],
       (floor(random() * (38 - 18 + 1) + 18))::int,
       CASE
           WHEN random() < 0.15 THEN 'goalkeeper'::player_position
           WHEN random() < 0.45 THEN 'defender'::player_position
           WHEN random() < 0.75 THEN 'midfielder'::player_position
           ELSE 'attacker'::player_position
           END,
       1000000
FROM generate_series(1, 500);