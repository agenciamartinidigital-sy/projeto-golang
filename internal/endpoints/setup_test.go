package endpoints

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	internalmock "projeto-golang/internal/test/internalMock"

	"github.com/go-chi/chi/v5"
)

var (
	service *internalmock.CampaignServiceMock
	handler = Handler{}
)

func setup() {

	service = new(internalmock.CampaignServiceMock)
	handler.CampaignService = service
}

func newReqAndRecord(method string, url string, valueParameter string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest(method, url, nil)
	rr := httptest.NewRecorder()
	return req, rr
}

func newHTTPTest(method string, url string, body interface{}) (*http.Request, *httptest.ResponseRecorder) {
	var buf bytes.Buffer
	if body != nil {
		json.NewEncoder(&buf).Encode(body)
	}
	req, _ := http.NewRequest(method, url, &buf)
	rr := httptest.NewRecorder()
	return req, rr
}

func addParamter(req *http.Request, keyParameter string, valueParameter string) *http.Request {
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add(keyParameter, valueParameter)
	return req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))
}

func addContext(req *http.Request, keyParameter string, valueParameter string) *http.Request {
	ctx := context.WithValue(req.Context(), keyParameter, valueParameter)
	return req.WithContext(ctx)

}
