package github_provider

import (
	"errors"
	"net/http"
	"testing"

	"github.com/iino123/golang-microservice-practice/src/api/domain/github"
	"github.com/iino123/golang-microservice-practice/src/api/restclient"
	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}
func TestCreateRepoError(t *testing.T) {
	restclient.StartMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:        errors.New("invalid client response"), // mockなのでerrorだったらなんでもいい
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid client response", err.Message)
}
