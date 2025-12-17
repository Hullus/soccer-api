package repo

import (
	"context"
	"errors"
	"soccer-api/internal/domain/responses"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MarketRepo struct {
	Pool *pgxpool.Pool
}

func (r MarketRepo) CreateListing(ctx context.Context, playerID, teamID int64, price int64) error {
	var status string
	checkQuery := `SELECT status FROM transfer_listings WHERE player_id = $1 AND status = 'active' LIMIT 1`
	err := r.Pool.QueryRow(ctx, checkQuery, playerID).Scan(&status)

	if err == nil {
		return errors.New("listing already exists")
	}

	insertQuery := `
		INSERT INTO transfer_listings (player_id, seller_team_id, asking_price_cents, status, listed_at)
		VALUES ($1, $2, $3, 'active', NOW())`

	_, err = r.Pool.Exec(ctx, insertQuery, playerID, teamID, price)
	return err
}

func (r MarketRepo) CancelListing(ctx context.Context, playerID int64) error {
	query := `
		UPDATE transfer_listings 
		SET status = 'cancelled' 
		WHERE player_id = $1 AND status = 'active'`

	_, err := r.Pool.Exec(ctx, query, playerID)
	return err
}

func (r MarketRepo) GetMarketListings(ctx context.Context) ([]responses.MarketListing, error) {
	query := `
		SELECT  p.first_name, p.last_name, p.country, p.position, tl.asking_price_cents
		FROM transfer_listings tl
		JOIN players p ON tl.player_id = p.id
		WHERE tl.status = 'active' 
		`

	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listings []responses.MarketListing
	for rows.Next() {
		var l responses.MarketListing
		if err := rows.Scan(&l.FirstName, &l.LastName, &l.Country, &l.Position, &l.AskingPriceCents); err != nil {
			return nil, err
		}
		listings = append(listings, l)
	}
	return listings, nil
}
