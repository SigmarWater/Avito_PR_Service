package pull_request

import "github.com/SigmarWater/Avito_PR_Service/internal/models"

func (s *Service) MergePullRequest(pullRequestId string) (*models.PullRequest, error) {
	pullRequest, err := s.MergePullRequest(pullRequestId)
	if err != nil {
		return nil, err
	}
	return pullRequest, nil
}
