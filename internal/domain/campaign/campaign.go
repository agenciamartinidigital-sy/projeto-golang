package campaign

import (
	"errors"
	internalerrors "projeto-golang/internal/internalErrors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Emails string `validate:"email"`
}

const (
	Pending string = "Pending"
	Started        = "Started"
	Done           = "Done"
)

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string
}

func NewCampaign(name, content string, emails []string) (*Campaign, error) {
	if emails == nil {
		return nil, errors.New("contacts is required with min 1")
	}
	if len(emails) == 0 {
		return nil, errors.New("contacts is required")
	}

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Emails = email
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(), // não pode ser nil
		Contacts:  contacts,
		Status:    "Pending",
	}

	err := internalerrors.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	}
	return nil, err
}
