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

func (r PlayerRepo) UpdatePlayer(ctx context.Context, playerID int64, firstName, lastName, country string) error {
	query := `UPDATE players SET first_name = $1, last_name = $2, country = $3, updated_at = NOW() WHERE id = $4`
	_, err := r.Pool.Exec(ctx, query, firstName, lastName, country, playerID)
	return err
}

func (r PlayerRepo) GetPlayerOwner(ctx context.Context, playerID int64) (int64, error) {
	var ownerID int64
	query := `SELECT t.owner_id FROM players p JOIN teams t ON p.team_id = t.id WHERE p.id = $1`
	err := r.Pool.QueryRow(ctx, query, playerID).Scan(&ownerID)
	return ownerID, err
}
