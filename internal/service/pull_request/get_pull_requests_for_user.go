package pull_request

import "github.com/SigmarWater/Avito_PR_Service/internal/models"

func (s *Service) GetPullRequestsForUser(userId string) (*models.UserWithPullRequests, error) {
	res, err := s.GetPullRequestsForUser(userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
