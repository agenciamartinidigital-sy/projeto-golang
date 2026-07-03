package endpoints

import (
	"projeto-golang/internal/domain/campaign"
)

type Handler struct {
	CampaignService campaign.Service
}
