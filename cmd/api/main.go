package main

import (
	"net/http"
	"projeto-golang/internal/domain/campaign"
	"projeto-golang/internal/endpoints"
	"projeto-golang/internal/infrastructure/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	PORT := ":3000"

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaingService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoints.Handler{
		CampaignService: &campaingService,
	}
	// handler.CampaingService = campaingService
	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaigns", endpoints.HandlerError(handler.CampaignGet))

	http.ListenAndServe(PORT, r)
}
