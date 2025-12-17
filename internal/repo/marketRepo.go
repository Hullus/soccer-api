package repo

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"soccer-api/internal/domain/responses"
	"strings"

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
		SELECT  tl.id, p.first_name, p.last_name, p.country, p.position, tl.asking_price_cents
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
		if err := rows.Scan(&l.ListingID, &l.FirstName, &l.LastName, &l.Country, &l.Position, &l.AskingPriceCents); err != nil {
			return nil, err
		}
		listings = append(listings, l)
	}
	return listings, nil
}

func (r MarketRepo) BuyPlayer(ctx context.Context, buyerUserID int64, listingID int64) error {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var buyerTeamID, buyerBudget int64
	err = tx.QueryRow(ctx, `SELECT id, budget_cents FROM teams WHERE owner_id = $1 FOR UPDATE`, buyerUserID).Scan(&buyerTeamID, &buyerBudget)
	if err != nil {
		return fmt.Errorf("failed to lock buyer team: %w", err)
	}

	var playerID, sellerTeamID, askingPrice int64
	var status string
	query := `
		SELECT tl.player_id, tl.seller_team_id, tl.asking_price_cents, tl.status 
		FROM transfer_listings tl
		JOIN players p ON tl.player_id = p.id
		JOIN teams t ON tl.seller_team_id = t.id
		WHERE tl.id = $1 FOR UPDATE OF tl, p, t`

	err = tx.QueryRow(ctx, query, listingID).Scan(&playerID, &sellerTeamID, &askingPrice, &status)
	if err != nil {
		return fmt.Errorf("listing not found: %w", err)
	}

	var validationErrors []string

	if status != "active" {
		validationErrors = append(validationErrors, "listing is no longer active")
	}
	if buyerTeamID == sellerTeamID {
		validationErrors = append(validationErrors, "cannot buy player from your own team")
	}
	if buyerBudget < askingPrice {
		validationErrors = append(validationErrors, "insufficient budget")
	}

	if len(validationErrors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(validationErrors, "; "))
	}

	_, err = tx.Exec(ctx, `UPDATE teams SET budget_cents = budget_cents - $1 WHERE id = $2`, askingPrice, buyerTeamID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, `UPDATE teams SET budget_cents = budget_cents + $1 WHERE id = $2`, askingPrice, sellerTeamID)
	if err != nil {
		return err
	}

	multiplier := 1.1 + rand.Float64()*(2.0-1.1)

	_, err = tx.Exec(ctx, `
    UPDATE players 
    SET team_id = $1, 
        market_value_cents = market_value_cents * $2,
        updated_at = NOW() 
    WHERE id = $3`, buyerTeamID, multiplier, playerID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `
    UPDATE transfer_listings
    SET status          = 'sold',
        sold_at         = NOW(),
        sold_to_team_id = $1
    WHERE id = $2`, buyerTeamID, listingID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `
    INSERT INTO transfers (
        player_id, 
        from_team_id, 
        to_team_id, 
        price_cents)
    VALUES ($1, $2, $3, $4)`,
		playerID, sellerTeamID, buyerTeamID, askingPrice)
	if err != nil {
		return fmt.Errorf("audit log insertion failed: %w", err)
	}

	return tx.Commit(ctx)
}
