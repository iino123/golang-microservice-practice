package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/iino123/golang-microservice-practice/src/api/domain/github"
	"github.com/iino123/golang-microservice-practice/src/api/restclient"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))
	response, err := restclient.Post(urlCreateRepo, request, headers)
	fmt.Println(response)
	fmt.Println(err)
	if err != nil {
		// Q: なぜprintfではだめか?
		// log.Printf("error when trying create new repo in github: %s", err.Error())
		log.Println(fmt.Sprintf("error when trying create new repo in github: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	// Q:なぜわざわざ[]byteに変換するのか？response.Bodyをそのまま扱うのではダメなのか?
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid response body",
		}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			log.Println(fmt.Sprintf("error when trying unmarshal GithubErrorResponse: %s", err.Error()))
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "error when trying unmarshal GithubErrorResponse",
			}
		}

		/*
		  githubのapiはstatusがresponseのbodyに含まれない。
		  --> bytesをUnMarshalしても、errResponseにStatusCodeはセットされていない
		  --> 下記のようにresponse.StatusCodeをセットする
		*/
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	// Q: json.Unmarshalの第二引数は、v interface{}となっているが、なぜポインタを渡す必要ある？
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying unmarshal CreateRepoResponse: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error when trying unmarshal CreateRepoResponse",
		}
	}

	return &result, nil
}
