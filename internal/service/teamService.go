package service

import (
	"context"
	"errors"
	"fmt"
	"soccer-api/internal/domain/responses"
	"soccer-api/internal/repo"
	"soccer-api/internal/util"
)

type TeamService struct {
	TeamRepo   repo.TeamRepo
	PlayerRepo repo.PlayerRepo
}

func (s TeamService) GetTeamInformation(ctx context.Context) (*responses.TeamInformationResponse, error) {
	userId := util.GetUserID(ctx)
	if userId == 0 {
		return nil, errors.New("unauthorized")
	}

	teamInfo, err := s.TeamRepo.GetTeamInformation(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("service.getTeamInformation: %w", err)
	}

	players, err := s.PlayerRepo.GetPlayersByTeam(ctx, teamInfo.ID)
	if err != nil {
		return nil, fmt.Errorf("service.getPlayersByTeam (ID: %d): %w", teamInfo.ID, err)
	}

	publicTeam := responses.TeamInformationPublic{
		Name:             teamInfo.Name,
		BudgetCents:      teamInfo.BudgetCents,
		Country:          teamInfo.Country,
		TotalMarketValue: teamInfo.TotalMarketValue,
	}

	return &responses.TeamInformationResponse{Team: publicTeam, Players: players}, nil
}

func (s TeamService) UpdateTeam(ctx context.Context, name, country string) error {
	userId := util.GetUserID(ctx)
	team, err := s.TeamRepo.GetTeamInformation(ctx, userId)
	if err != nil {
		return err
	}
	return s.TeamRepo.UpdateTeam(ctx, team.ID, name, country)
}

func (s TeamService) UpdatePlayer(ctx context.Context, playerID int64, firstName, lastName, country string) error {
	userId := util.GetUserID(ctx)
	ownerID, err := s.PlayerRepo.GetPlayerOwner(ctx, playerID)
	if err != nil {
		return err
	}
	if ownerID != userId {
		return errors.New("unauthorized: you do not own this player")
	}
	return s.PlayerRepo.UpdatePlayer(ctx, playerID, firstName, lastName, country)
}
