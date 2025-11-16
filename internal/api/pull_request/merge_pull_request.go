package pull_request

import (
	"context"
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

	ctx := context.Background()
	pullRequest, err := i.PullRequestService.MergePullRequest(ctx, req.PullRequestId)
	if err != nil {
		http.Error(w, "Invalid merge pull request", http.StatusInternalServerError)
	}

	render.JSON(w, r, pullRequest)
}
