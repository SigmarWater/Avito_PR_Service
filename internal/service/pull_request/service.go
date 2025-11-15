package pull_request

import "github.com/SigmarWater/Avito_PR_Service/internal/repository"

type Service struct {
	pullRequestRepository repository.PullRequestRepository
}

func NewService(pullRequestRepository repository.PullRequestRepository) *Service {
	return &Service{pullRequestRepository: pullRequestRepository}
}
