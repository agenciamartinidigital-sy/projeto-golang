package endpoints

import (
	"fmt"
	"projeto-golang/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	createdByExpectated = "teste@teste.com"
	body                = contract.NewCampaign{
		Name:    "teste",
		Content: "Hi, everyone",
		Emails:  []string{"teste@teste.com"},
	}
)

func Test_CampaignsPost_201(t *testing.T) {
	setup()

	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name == body.Name &&
			request.Content == body.Content &&
			request.CreatedBy == createdByExpectated {
			return true
		} else {
			return false
		}
	})).Return("34x", nil)
	req, rr := newHTTPTest("POST", "/", body)
	req = addContext(req, "email", createdByExpectated)

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(t, 201, status)
	assert.Nil(t, err)

}

func Test_CampaignsPost_Err(t *testing.T) {
	setup()

	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))

	req, rr := newHTTPTest("POST", "/", body)
	req = addContext(req, "email", createdByExpectated)
	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(t, err)
}
