package endpoints

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsStart_Return200(t *testing.T) {
	setup()
	campaignID := "xpto"
	service.On("Start", mock.MatchedBy(func(id string) bool {
		return id == campaignID
	})).Return(nil)
	req, rr := newReqAndRecord("PATCH", "/", "")
	req = addParamter(req, "id", campaignID)

	_, status, err := handler.CampaignStart(rr, req)

	assert.Equal(t, 200, status)
	assert.Nil(t, err)
}
func Test_CampaignsStart_ReturnErr(t *testing.T) {
	setup()

	errExpected := errors.New("Something wrong")
	service.On("Start", mock.Anything).Return(errExpected)
	req, rr := newReqAndRecord("PATCH", "/", "")

	_, _, err := handler.CampaignStart(rr, req)

	assert.Equal(t, errExpected, err)
}
