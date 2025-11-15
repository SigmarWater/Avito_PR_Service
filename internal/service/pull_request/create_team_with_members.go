package pull_request

import "github.com/SigmarWater/Avito_PR_Service/internal/models"

func (s *Service) CreateTeamWithMembers(req *models.Team) (*models.Team, error) {
	team, err := s.CreateTeamWithMembers(req)
	if err != nil {
		return nil, err
	}
	return team, nil
}
