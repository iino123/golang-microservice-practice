package domain

import (
	"fmt"
	"github.com/iino123/golang-microservice-practice/mvc/utils"
	"net/http"
)

// 定義部分で*Userという風にポインタ型にすることによって第一返り値としてnilを返すことを許容する
// Userとするとnilは返却できず、User{}を返す必要がある
func GetUser(userId int64) (*User, *utils.ApplicationError) {
	// 一行で書くことでuser変数をif内だけのローカル変数として扱える
	if user := users[userId]; user != nil {
		fmt.Println(user)
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}

// ここのカッコは何？
var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "ian", LastName: "Ken", Email: "iino@gmail.com"},
	}
)
