package services

import (
	"github.com/iino123/golang-microservice-practice/mvc/domain"
	"github.com/iino123/golang-microservice-practice/mvc/utils"
	"net/http"
)

type itemsService struct{}

// GetItemsを実装したitemsService(struct)を別パッケージでも使えるようにエクスポートするイメージ
// users_service.goを参照する
var (
	ItemsService itemsService
)

// なぜレシーバはポインタ型なのか? --> おそらくどっちでもいい。
func (i *itemsService) GetItems(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
