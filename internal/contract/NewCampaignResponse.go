package contract

type CampaignRespose struct {
	ID                 string
	Name               string
	Content            string
	Status             string
	AmountEmailsToSend int
	CreatedBy          string
}
