package services

import (
	"github.com/iino123/golang-microservice-practice/mvc/domain"
	"github.com/iino123/golang-microservice-practice/mvc/utils"
)

/*
今まではservice.GetUser(userId)のような形でアクセスしていた。
もう少し構造化したくて、service.UserService.GetUser(userId)という風にしたい。
↑(パッケージ名).(名前空間).(関数())
そんな時のやり方は以下の通り。
①空のstructを定義
②変数を定義
 - 別パッケージで使用するため大文字で始める
 - 上記の名前空間として使用する
③関数のレシーバを②で定義した変数とする
*/

// ①
type usersService struct{}

// type UsersService struct としてもいけそう... --> いけなかった
// TODO: ↑この理由
// ②
var (
	UsersService usersService
)

// なぜここのレシーバはポインタ型なのか?
// --> 多分どっちでも良いけど、迷ったらポインタ型でOK(値をコピーしないので)
// ③
func (u *usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
