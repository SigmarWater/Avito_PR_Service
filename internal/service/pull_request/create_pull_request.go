package pull_request

import (
	"context"
	"github.com/SigmarWater/Avito_PR_Service/internal/models"
)

func (s *Service) CreatePullRequest(ctx context.Context, req *models.CreatePullRequestRequest) (*models.PullRequest, error) {
	pullRequest, err := s.pullRequestRepository.CreatePullRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	return pullRequest, nil
}
