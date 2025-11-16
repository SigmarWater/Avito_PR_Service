package pull_request

import (
	"context"
	"github.com/SigmarWater/Avito_PR_Service/internal/models"
	"strconv"
)

func (s *Service) MergePullRequest(ctx context.Context, pullRequestId string) (*models.PullRequest, error) {
	pullRequestIdInt, _ := strconv.Atoi(pullRequestId)
	pullRequest, err := s.pullRequestRepository.MergePullRequest(ctx, pullRequestIdInt)
	if err != nil {
		return nil, err
	}
	return pullRequest, nil
}
