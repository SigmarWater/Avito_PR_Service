package pull_request

import "github.com/SigmarWater/Avito_PR_Service/internal/models"

func (s *Service) SetIsActive(userId string, isActive bool) (*models.User, error) {
	user, err := s.SetIsActive(userId, isActive)
	if err != nil {
		return nil, err
	}
	return user, nil
}
