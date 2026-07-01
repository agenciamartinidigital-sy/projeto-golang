package campaign

import (
	"projeto-golang/internal/contract"
	internalerrors "projeto-golang/internal/internalErrors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}
	// Precisa mascarar o erro
	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}
