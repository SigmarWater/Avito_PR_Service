package pull_request

import (
	"context"
	"github.com/SigmarWater/Avito_PR_Service/internal/models"
	"strconv"
)

func (s *Service) SetIsActive(ctx context.Context, userId string, isActive bool) (*models.User, error) {
	userIdInt, _ := strconv.Atoi(userId)
	user, err := s.pullRequestRepository.SetIsActive(ctx, userIdInt, isActive)
	if err != nil {
		return nil, err
	}
	return user, nil
}
