package pull_request

import (
	"github.com/SigmarWater/Avito_PR_Service/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

const (
	urlParamTeamName = "team_name"
)

func (i *Implementation) GetTeamWithMembers(w http.ResponseWriter, r *http.Request) {
	teamName := chi.URLParam(r, urlParamUserId)
	var teamWithMembers *models.Team
	teamWithMembers, _ = i.PullRequestService.GetTeamWithMembers(teamName)
	render.JSON(w, r, teamWithMembers)
}
