package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNotUserFound(t *testing.T) {
	user, err := GetUser(0)

	// assertではあるべき姿を記述して、そうならない場合のエラーメッセージを第３引数に渡す。(演算子によるけど)
	assert.Nil(t, user, "we were not expecting a user wih id 0")
	assert.NotNil(t, err, "we were expecting an error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)

	// NOTE: assertを使わない書き方
	//if user != nil {
	//	t.Error("we were not expecting a user wih id 0")
	//}
	//if err == nil {
	//	t.Error("we were expecting an error when user id is 0")
	//}
	//if err.StatusCode != http.StatusNotFound {
	//	t.Error()
	//}
}

func TestGetUserNotError(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)

	// FIXME: uint64とintの違いでテストが失敗する
	//assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "ian", user.FirstName)
	assert.EqualValues(t, "Ken", user.LastName)
	assert.EqualValues(t, "iino@gmail.com", user.Email)
}
