package pull_request

import (
	"context"
	"encoding/json"
	"github.com/SigmarWater/Avito_PR_Service/internal/models"
	"github.com/go-chi/render"
	"net/http"
)

func (i *Implementation) CreateTeamWithMembers(w http.ResponseWriter, r *http.Request) {
	var team models.Team

	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	ctx := context.Background()
	teamService, err := i.PullRequestService.CreateTeamWithMembers(ctx, &team)
	if err != nil {
		http.Error(w, "Invalid create team with members", http.StatusInternalServerError)
	}

	render.JSON(w, r, teamService)
}
