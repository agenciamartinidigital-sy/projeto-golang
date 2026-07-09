package endpoints

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"projeto-golang/internal/contract"
	internalmock "projeto-golang/internal/test/internalMock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup(body contract.NewCampaign, createdByExpectated string) (*http.Request, *httptest.ResponseRecorder) {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	ctx := context.WithValue(req.Context(), "email", createdByExpectated)
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()
	return req, rr
}

func Test_CampaignsPost_Save(t *testing.T) {
	assert := assert.New(t)
	createdByExpectated := "teste@teste.com"
	body := contract.NewCampaign{
		Name:    "teste",
		Content: "Hi, everyone",
		Emails:  []string{"teste@teste.com"},
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name == body.Name &&
			request.Content == body.Content &&
			request.CreatedBy == createdByExpectated {
			return true
		} else {
			return false
		}
	})).Return("34x", nil)
	handler := Handler{CampaignService: service}
	req, rr := setup(body, createdByExpectated)
	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(201, status)
	assert.Nil(err)

}

func Test_CampaignsPost_InformError(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "teste",
		Content: "Hi, everyone",
		Emails:  []string{"teste@teste.com"},
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}

	req, rr := setup(body, "teste@teste.com")
	// req, rr := setup(body, body.CreatedBy)

	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(err)
}
