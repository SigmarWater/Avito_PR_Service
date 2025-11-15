package pull_request

import "github.com/SigmarWater/Avito_PR_Service/internal/models"

func (s *Service) Reassign(req *models.ReassignRequest) (*models.ReassignResponse, error) {
	reassignResponse, err := s.Reassign(req)
	if err != nil {
		return nil, err
	}

	return reassignResponse, nil
}
