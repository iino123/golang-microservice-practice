package services

import (
	"github.com/iino123/golang-microservice-practice/mvc/domain"
	"github.com/iino123/golang-microservice-practice/mvc/utils"
)

// *domain.Userとしているのはなぜ？
// domain.*Userな感じするが、、、
/// -> あ、*(domain.User)ってことか
func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
