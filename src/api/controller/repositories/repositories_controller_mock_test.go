package repositories

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/iino123/golang-microservices/src/api/domain/repositories"
	"github.com/iino123/golang-microservices/src/api/services"
	"github.com/iino123/golang-microservices/src/api/utils/errors"
	"github.com/iino123/golang-microservices/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

var (
	funcCreateRepo func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
)

type repoServiceMock struct{}

func (s *repoServiceMock) CreateRepo(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	return funcCreateRepo(clientId, request)
}

func TestCreateRepoNoErrorMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return &repositories.CreateRepoResponse{
			Id:    321,
			Name:  "mocked service",
			Owner: "golang",
		}, nil
	}

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 321, result.Id)
	assert.EqualValues(t, "mocked service", result.Name)
	assert.EqualValues(t, "golang", result.Owner)
}

func TestCreateRepoErrorFromGithubMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid repository name", apiErr.Message())
}
