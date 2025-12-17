package repo

import (
	"context"
	"soccer-api/internal/domain/responses"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TeamRepo struct {
	Pool *pgxpool.Pool
}

func (r TeamRepo) GetTeamInformation(ctx context.Context, userId int64) (*responses.TeamInformation, error) {
	query := `SELECT id,
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
