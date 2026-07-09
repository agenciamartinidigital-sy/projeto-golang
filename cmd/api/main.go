package main

import (
	"fmt"
	"log"
	"net/http"
	"projeto-golang/internal/domain/campaign"
	"projeto-golang/internal/endpoints"
	"projeto-golang/internal/infrastructure/database"
	"projeto-golang/internal/infrastructure/database/mail"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	PORT := ":3000"

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	errMail := mail.SendEmail()
	if errMail != nil {
		log.Fatal("Error sending mail", errMail)
	}

	return

	route := chi.NewRouter()
	route.Use(middleware.RequestID)
	route.Use(middleware.ClientIPFromRemoteAddr)
	route.Use(middleware.Logger)
	route.Use(middleware.Recoverer)

	db := database.NewDB()

	campaingService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
	}
	handler := endpoints.Handler{
		CampaignService: &campaingService,
	}
	// handler.CampaingService = campaingService
	// Agrupamento de routes
	route.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)
		route.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})

		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetByID))
		r.Delete("/delete/{id}", endpoints.HandlerError(handler.CampaignDelete))

	})

	fmt.Println("Conexão estabelecida com sucesso")
	log.Fatal(http.ListenAndServe(PORT, route))
}
