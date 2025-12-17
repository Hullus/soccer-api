package service

import (
	"context"
	"errors"
	"soccer-api/internal/repo"
	"soccer-api/internal/util"
)

type MarketService struct {
	PlayerRepo repo.PlayerRepo
	MarketRepo repo.MarketRepo
	TeamRepo   repo.TeamRepo
}

func (s MarketService) CancelPlayerListing(ctx context.Context, playerID int64) error {
	userID := util.GetUserID(ctx)
	if userID == 0 {
		return errors.New("unauthorized")
	}

	ownerID, err := s.PlayerRepo.GetPlayerOwner(ctx, playerID)
	if err != nil {
		return err
	}
	if ownerID != userID {
		return errors.New("unauthorized: you do not own this player")
	}

	return s.MarketRepo.CancelListing(ctx, playerID)
}

func (s MarketService) ListPlayerOnMarket(ctx context.Context, playerID int64, price int64) error {
	userID := util.GetUserID(ctx)
	if userID == 0 {
		return errors.New("unauthorized")
	}

	ownerID, err := s.PlayerRepo.GetPlayerOwner(ctx, playerID)
	if err != nil {
		return err
	}
	if ownerID != userID {
		return errors.New("forbidden: you do not own this player")
	}

	team, err := s.TeamRepo.GetTeamInformation(ctx, userID)
	if err != nil {
		return err
	}

	return s.MarketRepo.CreateListing(ctx, playerID, team.ID, price)
}

func (s MarketService) BuyPlayer(ctx context.Context, listingID int64) error {
	userID := util.GetUserID(ctx)
	if userID == 0 {
		return errors.New("unauthorized")
	}

	return s.MarketRepo.BuyPlayer(ctx, userID, listingID)
}
