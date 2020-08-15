package app

import (
	"github.com/iino123/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.POST("/repository", repositories.CreateRepo)
}
