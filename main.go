package main

import (
	"github.com/gin-gonic/gin"
	"go-ddd-rest-api-sample/src/Infrastructures"
	"go-ddd-rest-api-sample/src/Presentations/Controllers/User"
)

func main() {
	db := Infrastructures.Init()
	defer db.Close()

	router := gin.Default()

	var controller User.UserController
	router.GET("/users/:userId", controller.GetUser)

	router.Run(":3000")
}
