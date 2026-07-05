package endpoints

import (
	"net/http"
	"net/http/httptest"
	"projeto-golang/internal/contract"
	internalmock "projeto-golang/internal/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsGet_Save(t *testing.T) {
	assert := assert.New(t)
	campaign := contract.CampaignRespose{
		ID:      "343",
		Name:    "Test",
		Content: "Hi!",
		Status:  "Pending",
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(&campaign, nil)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	response, status, _ := handler.CampaignGetByID(rr, req)

	assert.Equal(200, status)
	assert.Equal(campaign.ID, response.(*contract.CampaignRespose).ID)
	assert.Equal(campaign.Name, response.(*contract.CampaignRespose).Name)

}
