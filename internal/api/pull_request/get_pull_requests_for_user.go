package pull_request

import (
	"context"
	"github.com/SigmarWater/Avito_PR_Service/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

const (
	urlParamUserId = "user_id"
)

func (i *Implementation) GetPullRequestForUser(w http.ResponseWriter, r *http.Request) {
	var pullRequestForUser *models.UserWithPullRequests
	userID := chi.URLParam(r, urlParamUserId)
	ctx := context.Background()
	pullRequestForUser, _ = i.PullRequestService.GetPullRequestsForUser(ctx, userID)

	render.JSON(w, r, pullRequestForUser)
}
