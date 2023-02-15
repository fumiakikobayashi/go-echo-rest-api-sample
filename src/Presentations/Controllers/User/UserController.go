package User

import (
	"github.com/gin-gonic/gin"
	"go-ddd-rest-api-sample/src/Domains"
	"go-ddd-rest-api-sample/src/Infrastructures"
	"go-ddd-rest-api-sample/src/Infrastructures/Repositories"
	"go-ddd-rest-api-sample/src/UseCases/User"
	"net/http"
	"strconv"
)

type UserController struct{}

func (controller *UserController) GetUser(context *gin.Context) {
	userIdInt, _ := strconv.Atoi(context.Param("userId"))
	userId := Domains.UserId{Value: userIdInt}

	repository := Repositories.NewUserRepository(Infrastructures.GetDb())
	useCase := User.NewGetUserUseCase(repository)

	userDto := useCase.GetUser(userId)
	context.JSON(http.StatusOK, gin.H{
		"user": map[string]any{
			"userId":    userDto.Id,
			"firstName": userDto.FirstName,
			"lastName":  userDto.LastName,
		},
	})
}
