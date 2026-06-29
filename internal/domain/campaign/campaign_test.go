package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// testar o ponto de integração
// "t" é o contexto do teste
func Test_NewCampaign_CreateCampaign(t *testing.T) {
	// 3 "aa" Arrange (organização), Action (ação), Assert (reivindicar)

	// Arrange
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	// Action
	campaign := NewCampaign(name, content, contacts)

	// Assert
	// Equal = se os dois objetos são iguais
	assert.Equal(campaign.ID, campaign.ID)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}

func Test_NewCampaignIDIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

	// Obeto não é nil
	assert.NotNil(campaign.ID)
}

func TestNewCampaign_CreatedOnIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

	// Obeto não é nil
	assert.NotNil(campaign.CreatedOn)
}
