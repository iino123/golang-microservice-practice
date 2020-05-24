package domain

import (
	"fmt"
	"github.com/iino123/golang-microservice-practice/mvc/utils"
	"log"
	"net/http"
)

//  domain.UserDao.GetUser(userId)のような形で呼び出すためのもの
// 	userDaoはこのinterfaceを満たすもので置き換えることができるようになる(users_service_test.goで使用)
var UserDao usersServiceInterface

type userDao struct{}

// ファイルの初期化時点でUserDaoを指定
// users_service_test.goでは、DBにアクセスしないようにモックを代入している
func init() {
	UserDao = &userDao{}
}

// users_service_test.goでdbにアクセスせずにモックを使ってテストするためにinterfaceを定義
type usersServiceInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

// 定義部分で*Userという風にポインタ型にすることによって第一返り値としてnilを返すことを許容する
// Userとするとnilは返却できず、User{}を返す必要がある
func (*userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("accessing the database!!!!!")

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

// ここのカッコは何？ --> 変数を複数定義するためのカッコ(まあまだ１つしかないけどね)
var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "ian", LastName: "Ken", Email: "iino@gmail.com"},
	}
)
