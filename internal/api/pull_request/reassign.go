package pull_request

import (
	"encoding/json"
	"github.com/SigmarWater/Avito_PR_Service/internal/models"
	"github.com/go-chi/render"
	"net/http"
)

func (i *Implementation) Reassign(w http.ResponseWriter, r *http.Request) {

	var req models.ReassignRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	resp, err := i.PullRequestService.Reassign(&req)
	if err != nil {
		http.Error(w, "Invalid reassign", http.StatusInternalServerError)
	}

	render.JSON(w, r, resp)
}
