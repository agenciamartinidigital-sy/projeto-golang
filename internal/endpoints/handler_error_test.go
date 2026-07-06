package endpoints

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	internalerrors "projeto-golang/internal/internalErrors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandlerError(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.ErrInternal
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())
}

func Test_HandlerErrorDomain(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("domain error")
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "domain error")
}

func Test_ObjectAndStatus(t *testing.T) {
	assert := assert.New(t)
	type bodyForTest struct {
		ID int `json:"id"`
	}
	objExprected := bodyForTest{ID: 1}
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return objExprected, 201, nil
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)
	objReturned := bodyForTest{}
	err := json.Unmarshal(res.Body.Bytes(), &objReturned)
	assert.NoError(err)
	assert.Equal(objExprected, objReturned)
}
