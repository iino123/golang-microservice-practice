package app

import (
	"github.com/iino123/golang-microservice-practice/mvc/controllers"
)

func mapUrls() {
	//http.HandleFunc("users/", controllers.GetUser)
	router.GET("/users/:user_id", controllers.GetUser)
}
