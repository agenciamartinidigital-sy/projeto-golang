package endpoints

import (
	"net/http"
	"projeto-golang/internal/contract"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaign
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		return nil, http.StatusBadRequest, err
	}
	id, err := h.CampaignService.Create(request)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return map[string]string{"id": id}, 201, err
}
