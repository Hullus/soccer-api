package repo

import (
	"context"
	"fmt"
	"math/rand/v2"
	"soccer-api/internal/domain/responses"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TeamRepo struct {
	Pool *pgxpool.Pool
}

func (r TeamRepo) GetTeamInformation(ctx context.Context, userId int64) (*responses.TeamInformation, error) {
	query := `
        SELECT id,
               name,
               budget_cents,
               country,
               owner_id,
               (SELECT SUM(market_value_cents) FROM players WHERE team_id = teams.id) AS total_market_value
        FROM teams
        WHERE owner_id = $1;
`

	var t responses.TeamInformation
	err := r.Pool.QueryRow(ctx, query, userId).Scan(
		&t.ID,
		&t.Name,
		&t.BudgetCents,
		&t.Country,
		&t.OwnerID,
		&t.TotalMarketValue,
	)

	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r TeamRepo) UpdateTeam(ctx context.Context, teamID int64, name, country string) error {
	query := `UPDATE teams SET name = $1, country = $2 WHERE id = $3`
	_, err := r.Pool.Exec(ctx, query, name, country, teamID)
	return err
}

func (r TeamRepo) AssignNewTeam(ctx context.Context, userID int64, email string) error {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	teams := []string{"Real Madrid", "Manchester City", "Bayern Munich", "Paris Saint-Germain", "Liverpool", "Inter Milan", "Arsenal", "Barcelona", "Borussia Dortmund", "Juventus", "Atletico Madrid", "Bayer Leverkusen", "AC Milan"}
	countries := []string{"England", "Spain", "Germany", "Georgia", "France", "Italy", "Brazil", "Argentina", "Netherlands", "Portugal", "Belgium"}
	teamCountry := countries[rand.IntN(len(countries))]
	teamName := teams[rand.IntN(len(teams))]
	var teamID int64
	teamQuery := `INSERT INTO teams (name, country, owner_id) VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRow(ctx, teamQuery, teamName, teamCountry, userID).Scan(&teamID)
	if err != nil {
		return fmt.Errorf("failed to create team: %w", err)
	}

	requiredPositions := map[string]int{
		"goalkeeper": 3,
		"defender":   6,
		"midfielder": 6,
		"attacker":   5,
	}

	for pos, count := range requiredPositions {
		selectQuery := `
			SELECT id FROM players 
			WHERE team_id IS NULL AND position = $1::player_position
			ORDER BY RANDOM() 
			LIMIT $2
			FOR UPDATE SKIP LOCKED
		`

		rows, err := tx.Query(ctx, selectQuery, pos, count)
		if err != nil {
			return fmt.Errorf("failed to query free agents for %s: %w", pos, err)
		}

		var playerIDs []int64
		for rows.Next() {
			var pid int64
			if err := rows.Scan(&pid); err != nil {
				return err
			}
			playerIDs = append(playerIDs, pid)
		}
		rows.Close()

		if len(playerIDs) < count {
			return fmt.Errorf("insufficient free agents for position %s: need %d, found %d", pos, count, len(playerIDs))
		}

		updateQuery := `UPDATE players SET team_id = $1 WHERE id = ANY($2)`
		_, err = tx.Exec(ctx, updateQuery, teamID, playerIDs)
		if err != nil {
			return fmt.Errorf("failed to assign players: %w", err)
		}
	}

	return tx.Commit(ctx)
}
