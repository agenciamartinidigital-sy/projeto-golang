package campaign

import (
	"time"

	"github.com/rs/xid"
)

// Validação por meio de anotação
type Contact struct {
	Emails string `validate:"email"`
}

// Validação por meio de anotação
type Campaign struct {
	ID        string    `validate:"required`
	Name      string    `validate:"min=5,max=24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
}

func NewCampaign(name, content string, emails []string) (*Campaign, error) {

	// if name == "" {
	// 	return nil, errors.New("name is required")
	// } else if content == "" {
	// 	return nil, errors.New("content is required")
	// } else if len(emails) == 0 {
	// 	return nil, errors.New("contacts is required")
	// }

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Emails = email
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(), // não pode ser nil
		Contacts:  contacts,
	}, nil
}
