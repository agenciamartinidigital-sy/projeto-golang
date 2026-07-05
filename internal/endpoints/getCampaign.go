package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGetByID(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	// GET /campaigns/{id}
	id := chi.URLParam(r, "id")

	campaign, err := h.CampaignService.GetBy(id)
	return campaign, 200, err
}
