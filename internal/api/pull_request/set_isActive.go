package pull_request

import (
	"encoding/json"
	"github.com/SigmarWater/Avito_PR_Service/internal/models"
	"github.com/go-chi/render"
	"net/http"
)

func (i *Implementation) SetIsActive(w http.ResponseWriter, r *http.Request) {
	var req models.SetIsActiveRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	user, err := i.PullRequestService.SetIsActive(req.UserId, req.IsActive)
	if err != nil {
		http.Error(w, "Invalid set is active", http.StatusInternalServerError)
	}

	render.JSON(w, r, user)
}
