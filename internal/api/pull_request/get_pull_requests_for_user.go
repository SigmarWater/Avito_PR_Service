package pull_request

import (
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
	pullRequestForUser, _ = i.PullRequestService.GetPullRequestsForUser(userID)

	render.JSON(w, r, pullRequestForUser)
}
