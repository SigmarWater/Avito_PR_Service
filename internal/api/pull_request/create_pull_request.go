package pull_request

import (
	"encoding/json"
	"github.com/SigmarWater/Avito_PR_Service/internal/models"
	"github.com/go-chi/render"
	"net/http"
)

func (i *Implementation) CreatePullRequest(w http.ResponseWriter, r *http.Request) {
	var pullRequest models.CreatePullRequestRequest
	if err := json.NewDecoder(r.Body).Decode(&pullRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	pullRequestService, err := i.PullRequestService.CreatePullRequest(&pullRequest)
	if err != nil {
		http.Error(w, "Invalid create pull request", http.StatusInternalServerError)
	}

	render.JSON(w, r, pullRequestService)
}
