package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"soccer-api/internal/domain/responses"
	"soccer-api/internal/repo"
	"soccer-api/internal/util"
)

type TeamService struct {
	teamRepo   repo.TeamRepo
	playerRepo repo.PlayerRepo
}

func (r TeamService) GetTeamInformation(ctx context.Context) (*responses.TeamInformationResponse, error) {
	userId := util.GetUserID(ctx)
	if userId == 0 {
		return nil, nil //TODO CHANGE THIS
	}

	teamInfo, err := r.teamRepo.GetTeamInformation(ctx, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil //TODO CHANGE THIS
		}
		return nil, fmt.Errorf("service.getTeamInformation: %w", err)
	}

	players, err := r.playerRepo.GetPlayersByTeam(ctx, teamInfo.ID)
	if err != nil {
		return nil, fmt.Errorf("service.getPlayersByTeam (ID: %d): %w", teamInfo.ID, err)
	}

	return &responses.TeamInformationResponse{Team: *teamInfo, Players: players}, nil
}
