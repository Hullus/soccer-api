package repo

import (
	"context"
	"fmt"
	"soccer-api/internal/domain/responses"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PlayerRepo struct {
	Pool *pgxpool.Pool
}

func (r PlayerRepo) GetPlayersByTeam(ctx context.Context, teamID int64) ([]responses.Player, error) {
	query := `
        SELECT 
            first_name,
            last_name,
            country,
            age,
            market_value_cents,
            position
        FROM players
        WHERE team_id = $1;
    `

	rows, err := r.Pool.Query(ctx, query, teamID)
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	var players []responses.Player
	for rows.Next() {
		var p responses.Player
		err := rows.Scan(
			&p.FirstName,
			&p.LastName,
			&p.Country,
			&p.Age,
			&p.MarketValueCents,
			&p.Position,
		)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		players = append(players, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return players, nil
}
