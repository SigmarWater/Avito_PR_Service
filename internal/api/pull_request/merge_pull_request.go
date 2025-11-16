package pull_request

import (
	"encoding/json"
	"github.com/SigmarWater/Avito_PR_Service/internal/models"
	"github.com/go-chi/render"
	"net/http"
)

func (i *Implementation) MergePullRequest(w http.ResponseWriter, r *http.Request) {

	var req models.MergePullRequestRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	pullRequest, err := i.PullRequestService.MergePullRequest(req.PullRequestId)
	if err != nil {
		http.Error(w, "Invalid merge pull request", http.StatusInternalServerError)
	}

	render.JSON(w, r, pullRequest)
}
