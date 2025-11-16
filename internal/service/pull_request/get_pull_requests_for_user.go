package pull_request

import (
	"context"
	"github.com/SigmarWater/Avito_PR_Service/internal/models"
	"strconv"
)

func (s *Service) GetPullRequestsForUser(ctx context.Context, userId string) (*models.UserWithPullRequests, error) {
	userIdInt, _ := strconv.Atoi(userId)
	res, err := s.pullRequestRepository.GetPullRequestsForUser(ctx, userIdInt)
	if err != nil {
		return nil, err
	}

	return res, nil
}
