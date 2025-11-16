package pull_request

import (
	"context"
	"github.com/SigmarWater/Avito_PR_Service/internal/models"
)

func (s *Service) CreateTeamWithMembers(ctx context.Context, req *models.Team) (*models.Team, error) {
	team, err := s.pullRequestRepository.CreateTeamWithMembers(ctx, req)
	if err != nil {
		return nil, err
	}
	return team, nil
}
