package pull_request

import "github.com/SigmarWater/Avito_PR_Service/internal/service"

type Implementation struct {
	PullRequestService service.PullRequestService
}

func NewImplementation(pullRequestService service.PullRequestService) *Implementation {
	return &Implementation{PullRequestService: pullRequestService}
}
