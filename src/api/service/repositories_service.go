package services

import (
	"github.com/iino123/golang-microservices/src/api/config"
	"github.com/iino123/golang-microservices/src/api/domain/github"
	"github.com/iino123/golang-microservices/src/api/domain/repositories"
	"github.com/iino123/golang-microservices/src/api/log/option_b"
	"github.com/iino123/golang-microservices/src/api/providers/github_provider"
	"github.com/iino123/golang-microservices/src/api/utils/errors"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(clientId string, input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}
	// option_b.Info("about to send request to external api",
	// 	option_b.Field("client_id", clientId),
	// 	option_b.Field("status", "pending"),
	// 	option_b.Field("authenticated", clientId != ""))
	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		option_b.Error("response obtained from external api", err,
			option_b.Field("client_id", clientId),
			option_b.Field("status", "error"),
			option_b.Field("authenticated", clientId != ""))
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	option_b.Info("response obtained from external api",
		option_b.Field("client_id", clientId),
		option_b.Field("status", "success"),
		option_b.Field("authenticated", clientId != ""))

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}
