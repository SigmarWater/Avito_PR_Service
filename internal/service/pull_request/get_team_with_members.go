package pull_request

import (
	"context"
	"github.com/SigmarWater/Avito_PR_Service/internal/models"
)

func (s *Service) GetTeamWithMembers(ctx context.Context, teamName string) (*models.Team, error) {
	team, err := s.pullRequestRepository.GetTeamWithMembers(ctx, teamName)
	if err != nil {
		return nil, err
	}
	return team, nil
}
