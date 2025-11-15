package pull_request

import "github.com/SigmarWater/Avito_PR_Service/internal/models"

func (s *Service) GetTeamWithMembers(teamName string) (*models.Team, error) {
	team, err := s.GetTeamWithMembers(teamName)
	if err != nil {
		return nil, err
	}
	return team, nil
}
