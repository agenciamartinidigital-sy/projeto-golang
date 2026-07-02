package campaign

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"projeto-golang/internal/contract"
	internalerrors "projeto-golang/internal/internalErrors"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaign, error) {
	// args := r.Called(campaign)
	return nil, nil
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body Hi!",
		Emails:  []string{"teste1@test.com"},
	}

	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomaininError(t *testing.T) {
	assert := assert.New(t)
	// newCampaign.Name = "" -> nunca criar dependência de teste
	// _, err := service.Create(newCampaign)
	_, err := service.Create(contract.NewCampaign{})

	assert.False(errors.Is(err, internalerrors.ErrInternal))
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service.Repository = repositoryMock
	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("Error to save"))

	service.Repository = repositoryMock
	_, err := service.Create(newCampaign)

	assert.ErrorIs(err, internalerrors.ErrInternal)
	repositoryMock.AssertExpectations(t)

}
