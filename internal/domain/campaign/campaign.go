package campaign

import (
	"errors"
	internalerrors "projeto-golang/internal/internalErrors"
	"time"

	"github.com/rs/xid"
)

const (
	Pending  string = "Pending"
	Updated         = "Updated"
	Canceled        = "Canceled"
	Started         = "Started"
	Done            = "Done"
	Fail            = "Fail"
	Deleted         = "Deleted"
)

type Contact struct {
	ID         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignID string `gorm:"size:50"`
}

type Campaign struct {
	ID        string    `json:"id" validate:"required" gorm:"size:50;not null"`
	Name      string    `json:"name" validate:"min=5,max=24" gorm:"size:100;not null"`
	CreatedOn time.Time `json:"createdOn" validate:"required" gorm:"not null"`
	UpdatedOn time.Time `json:"updatedOn"`
	Content   string    `json:"content" validate:"min=5,max=1024" gorm:"size:1024;not null"`
	Contacts  []Contact `json:"contacts" validate:"min=1,dive"`
	Status    string    `json:"status" gorm:"size:20;not null"`
	CreatedBy string    `json:"createdBy" validate:"email" gorm:"size:50;not null"`
}

func NewCampaign(name, content string, emails []string, createdby string) (*Campaign, error) {
	if emails == nil {
		return nil, errors.New("contacts is required with min 1")
	}
	if len(emails) == 0 {
		return nil, errors.New("contacts is required")
	}

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
		contacts[index].ID = xid.New().String()
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
		Status:    Pending,
		CreatedBy: createdby,
	}

	err := internalerrors.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	}
	return nil, err
}

func (c *Campaign) Create() {
	c.Status = Started
}

func (c *Campaign) Fail() {
	c.Status = Fail
}

func (c *Campaign) Done() {
	c.Status = Done
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
}

func (c *Campaign) Update() {
	c.Status = Updated
}

func (c *Campaign) Delete() {
	c.Status = Done
}

func (c *Campaign) Start() error {
	if c.Status != Pending {
		return errors.New("Apenas as camapanhas pendentes podem ser iniciadas")
	}
	c.Status = Started
	return nil
}

func (c *Campaign) Finished() error {
	if c.Status != Started {
		return errors.New("Campaign Finalizada")
	}
	c.Status = Done
	return nil
}
